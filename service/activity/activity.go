package activity

import (
	"skyshi-technical-test/model/domain/activity"
	response "skyshi-technical-test/model/web"
	request_body "skyshi-technical-test/model/web/request_body/activity"
)

type activityService struct {
	repository activity.Repository
}

// Create implements activity.Service
func (as *activityService) Create(body *request_body.Create) response.Response {
	// cal related repository method for creating activity
	activity, err := as.repository.Create(body)

	// if error occur when persisting activity data
	if err != nil {
		return response.Response{
			Status:  "Failed",
			Message: "Server Error",
		}
	}

	return response.Response{
		Status:  "Success",
		Message: "Success",
		Data:    activity,
	}
}

// Delete implements activity.Service
func (as *activityService) Delete(body *request_body.Delete) response.Response {
	// fetch selected activity
	activity := as.repository.GetOne(body.ID)

	// if activity not found
	if activity.ID == 0 {
		return response.Response{
			Status:  "Failed",
			Message: "Activity not found",
		}
	}

	// if activity already deleted
	if activity.DeletedAt != "" && activity.ID != 0 {
		return response.Response{
			Status:  "Failed",
			Message: "Activity already deleted",
		}
	}

	// if activity can be deleted
	if err := as.repository.Delete(body.ID); err != nil {
		return response.Response{
			Message: "Server Error",
			Status:  "Failed",
		}
	}

	return response.Response{
		Status:  "Success",
		Message: "Success",
	}
}

// GetAll implements activity.Service
func (as *activityService) GetAll() response.Response {
	// call related respository method for qerying
	activities := as.repository.GetAll()

	// if activities is empty
	if len(activities) == 0 {
		return response.Response{
			Status:  "Success",
			Message: "There is no activity yet",
		}
	}

	return response.Response{
		Status:  "Success",
		Message: "Success",
		Data:    activities,
	}
}

// GetOne implements activity.Service
func (as *activityService) GetOne(body *request_body.GetOne) response.Response {
	// call related respository method for qerying
	activity := as.repository.GetOne(body.ID)

	// if activities is empty
	if activity.ID == 0 {
		return response.Response{
			Status:  "Success",
			Message: "Activity not found",
		}
	}

	return response.Response{
		Status:  "Success",
		Message: "Success",
		Data:    activity,
	}
}

func ActivityService(repo *activity.Repository) activity.Service {
	return &activityService{repository: *repo}
}
