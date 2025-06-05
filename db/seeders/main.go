package seeders

import (
	"fmt"

	"gorm.io/gorm"
)

type Seeder interface {
	Seed(db *gorm.DB) error
}

type DatabaseSeeder struct {
	seeders []Seeder
}

func (ds *DatabaseSeeder) AddSeeder(seeder Seeder) {
	ds.seeders = append(ds.seeders, seeder)
}

func (ds *DatabaseSeeder) Run(db *gorm.DB) error {
	fmt.Println("Starting database seeding...")

	for _, seeder := range ds.seeders {
		seederName := fmt.Sprintf("%T", seeder)
		fmt.Printf("Running %s...\n", seederName)

		if err := seeder.Seed(db); err != nil {
			return fmt.Errorf("seeder %s failed: %w", seederName, err)
		}

		fmt.Printf("%s completed successfully\n", seederName)
	}

	fmt.Println("Database seeding completed!")
	return nil
}
