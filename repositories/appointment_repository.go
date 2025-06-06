package repositories

import (
	"time"
	"totesbackend/models"

	"gorm.io/gorm"
)

type AppointmentRepository struct {
	DB *gorm.DB
}

func NewAppointmentRepository(db *gorm.DB) *AppointmentRepository {
	return &AppointmentRepository{DB: db}
}

func (r *AppointmentRepository) GetAppointmentByID(id int) (*models.Appointment, error) {
	var appointment models.Appointment
	err := r.DB.First(&appointment, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &appointment, nil
}

func (r *AppointmentRepository) GetAllAppointments() ([]models.Appointment, error) {
	var appointments []models.Appointment
	err := r.DB.Find(&appointments).Error
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

func (r *AppointmentRepository) SearchAppointmentsByState(state bool) ([]models.Appointment, error) {
	var appointments []models.Appointment
	err := r.DB.Where("state = ?", state).Find(&appointments).Error
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

func (r *AppointmentRepository) GetAppointmentsByCustomerID(customerID int) ([]models.Appointment, error) {
	var appointments []models.Appointment
	err := r.DB.Where("customer_id = ?", customerID).Find(&appointments).Error
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

func (r *AppointmentRepository) CreateAppointment(appointment *models.Appointment) (*models.Appointment, error) {
	if err := r.DB.Create(appointment).Error; err != nil {
		return nil, err
	}
	return appointment, nil
}

func (r *AppointmentRepository) UpdateAppointment(appointment *models.Appointment) error {
	if err := r.DB.Save(appointment).Error; err != nil {
		return err
	}
	return nil
}

func (r *AppointmentRepository) SearchAppointmentsByID(query string) ([]models.Appointment, error) {
	var appointments []models.Appointment
	err := r.DB.Where("CAST(id AS TEXT) LIKE ?", query+"%").Find(&appointments).Error
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

func (r *AppointmentRepository) SearchAppointmentsByCustomerID(query string) ([]models.Appointment, error) {
	var appointments []models.Appointment
	err := r.DB.Where("CAST(customer_id AS TEXT) LIKE ?", query+"%").Find(&appointments).Error
	if err != nil {
		return nil, err
	}
	return appointments, nil
}

func (r *AppointmentRepository) GetAppointmentByCustomerIDAndDate(customerID int, dateTime time.Time) (*models.Appointment, error) {
	var appointment models.Appointment
	err := r.DB.Where("customer_id = ? AND date_time = ?", customerID, dateTime).First(&appointment).Error
	if err != nil {
		return nil, err
	}
	return &appointment, nil
}

func (r *AppointmentRepository) CountAppointmentsAtDateTime(dateTime time.Time) (int64, error) {
	var count int64
	err := r.DB.Model(&models.Appointment{}).
		Where("date_time = ?", dateTime).
		Count(&count).Error
	return count, err
}

func (r *AppointmentRepository) DeleteAppointmentByID(id int) error {
	result := r.DB.Delete(&models.Appointment{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (r *AppointmentRepository) CountAppointmentsByHourOnDate(date time.Time) ([]int, error) {
	counts := make([]int, 9) // from 9:00 to 17:00

	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 9, 0, 0, 0, date.Location())
	endOfDay := time.Date(date.Year(), date.Month(), date.Day(), 17, 59, 59, 0, date.Location())

	var appointments []models.Appointment
	err := r.DB.Where("date_time BETWEEN ? AND ?", startOfDay, endOfDay).Find(&appointments).Error
	if err != nil {
		return nil, err
	}

	for _, appointment := range appointments {
		hour := appointment.DateTime.Hour()
		if hour >= 9 && hour <= 17 {
			index := hour - 9
			counts[index]++
		}
	}

	return counts, nil
}
