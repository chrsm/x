package testutil

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-pg/pg"
	"gopkg.in/yaml.v2"
)

type Fixture map[string]interface{}
type FixtureSet map[string]Fixture
type Fixtures struct {
	f map[string]Fixture

	// map names to insert ID
	ids map[string]int64
}

func newFixtures() *Fixtures {
	return &Fixtures{make(map[string]Fixture), make(map[string]int64)}
}

func (f *Fixtures) Get(name string) Fixture {
	return f.f[name]
}

func (f *Fixtures) ID(name string) int64 {
	return f.ids[name]
}

func LoadFixture(path string) (Fixture, error) {
	fi, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	dec := yaml.NewDecoder(fi)

	set := make(Fixture)
	if err := dec.Decode(set); err != nil {
		return nil, err
	}

	return set, nil
}

// LoadFixtures enumerates .yml files in a directory and loads them as fixtures.
func LoadFixtures(path string) (FixtureSet, error) {
	dir, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	files, err := dir.Readdirnames(0)
	if err != nil {
		return nil, err
	}

	m := make(FixtureSet)
	for i := range files {
		fixture, err := LoadFixture(path + "/" + files[i])
		if err != nil {
			return nil, err
		}

		m[files[i]] = fixture
	}

	return m, nil
}

// InsertFixtures inserts a FixtureSet into a *pg.DB.
func InsertFixtures(db *pg.DB, fs FixtureSet) (*Fixtures, error) {
	type colres struct {
		ColumnName string
		DataType   string
	}
	type idret struct {
		ID int64
	}

	ret := newFixtures()
	for fname := range fs {
		// this is ghetto.. :-)
		spl := strings.Split(fname, ".")
		if len(spl) != 2 {
			return nil, errors.New(fmt.Sprintf("can't create db fixtures for %s", fname))
		}

		table := spl[0]

		var mdl []*colres
		cols := make(map[string]bool)
		db.Query(&mdl, `SELECT column_name FROM INFORMATION_SCHEMA.COLUMNS WHERE table_name = ?`, table)
		for i := range mdl {
			cols[mdl[i].ColumnName] = true
		}

		db.Exec("TRUNCATE " + table)

		// now we know what this table wants. for each fixture in this file, try to insert it,
		// keeping track of the name of the fixture entry so we can reference it.
		for i := range fs[fname] {
			ent := fs[fname][i]

			var col, placeholder []string
			var val []interface{}

			realent, ok := ent.(map[interface{}]interface{})
			if !ok {
				log.Printf("%#v is not map[string]iface", ent)
				continue
			}

			for j := range realent {
				col = append(col, j.(string))
				val = append(val, realent[j])
				placeholder = append(placeholder, "?")
			}

			cols := strings.Join(col, ", ")
			placeholders := strings.Join(placeholder, ", ")
			_, _, _ = val, cols, placeholders

			ider := &idret{}
			_, err := db.QueryOne(ider,
				fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s) RETURNING id", table, cols, placeholders),
				val...,
			)

			if err != nil {
				panic(err)
			}

			intint := make(map[string]interface{})
			for i := range realent {
				intint[i.(string)] = realent[i]
			}
			ret.f[i] = Fixture(intint)
			ret.ids[i] = ider.ID
		}

		// update the sequences
		db.Exec(fmt.Sprintf("SELECT setval('%s_id_seq', (SELECT MAX(id) FROM %s));", table, table))
	}

	return ret, nil
}
