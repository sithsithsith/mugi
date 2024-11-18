package services

type MigrationService struct {
	CognitoService  *CognitoService
	DatabaseService *DatabaseService
}

func NewMigrationService(cs *CognitoService, ds *DatabaseService) *MigrationService {
	return &MigrationService{
		CognitoService:  cs,
		DatabaseService: ds,
	}
}

func (ms *MigrationService) MigrateUsers() error {
	users, err := ms.DatabaseService.FetchUsers()
	if err != nil {
		return err
	}

	for _, user := range users {
		err := ms.CognitoService.SignUp(user["phone_number"].(string), "temporaryPassword123")
		if err != nil {
			return err
		}
	}

	return nil
}
