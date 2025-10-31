package DBAdapter

import (
	"backend/DBAdapter/Models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Adapter struct {
	db *gorm.DB
}

func InitAdapter(host string, port string, user string, password string, dbname string) Adapter {
	query := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(query), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	adapter := Adapter{
		db: db,
	}
	adapter.migrate()
	return adapter
}

func (Adapter *Adapter) Blank() {

}

func (Adapter *Adapter) Close() {
	db, err := Adapter.db.DB()
	if err != nil {
		log.Fatal(err)
	}
	db.Close()
}

func (adapter *Adapter) migrate() {
	var err error
	err = Models.MigrateBaseTpuUsers(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating BaseTpuUsers\nError:\n%+v", err)
	}
	err = Models.MigrateEducationTypes(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating EducationTypes\nError:\n%+v", err)
	}
	err = Models.MigrateEngineeringSchools(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating EngineeringSchools\nError:\n%+v", err)
	}
	err = Models.MigrateDepartments(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating Departments\nError:\n%+v", err)
	}
	err = Models.MigrateTutors(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating Tutors\nError:\n%+v", err)
	}
	err = Models.MigrateRoles(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating Roles\nError:\n%+v", err)
	}
	err = Models.MigratePlanners(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating Planners\nError:\n%+v", err)
	}
	err = Models.MigrateQualificationLevels(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating QualificationLevels\nError:\n%+v", err)
	}
	err = Models.MigrateGroups(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating Groups\nError:\n%+v", err)
	}
	err = Models.MigrateSpecializations(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating Specializations\nError:\n%+v", err)
	}
	err = Models.MigrateStudents(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating Students\nError:\n%+v", err)
	}
	err = Models.MigrateSubjectDirectionRelations(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating SubjectDirectionRelations\nError:\n%+v", err)
	}
	err = Models.MigrateSubjectSpecializationRelations(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating SubjectSpecializationRelations\nError:\n%+v", err)
	}
	err = Models.MigrateSubjectTutorRelations(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating SubjectTutorRelations\nError:\n%+v", err)
	}
	err = Models.MigrateSubjects(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating Subjects\nError:\n%+v", err)
	}
	err = Models.MigrateTrainingDirections(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating TrainingDirections\nError:\n%+v", err)
	}
	err = Models.MigrateTutorDepartmentRelations(adapter.db)
	if err != nil {
		log.Fatalf("Error migrating TutorDepartmentRelations\nError:\n%+v", err)
	}

	log.Println("Successful Migration")
}
