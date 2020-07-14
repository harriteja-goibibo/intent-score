package real_time_data

import (
	"context"
	"github.com/goibibo/intent-score/internal/core"
	"github.com/goibibo/intent-score/internal/core/common"
	timeHelpers "github.com/goibibo/intent-score/internal/helpers/time"
	"github.com/goibibo/intent-score/pkg/api/grpc/manthan"
	"log"
	"strconv"
	"time"
)

func SaveRealTimeData(ctx context.Context, request *manthan.ManthanRealTimeData, repo core.ScoreRepository) (err error) {
	// 1. Check new/old to this system if old, update the normal flow, otherwise get data from redshift and update the old transaction
	userHistory, _ := repo.GetUserData(ctx, request.UserId, request.Vertical.String())
	if len(userHistory.UpdatedTs) == 0 {
		// TODO: Get User Transaction Details from Redshift and update the data
		// Initially data won't be there in this service So, First time
		err = repo.SetUserData(ctx, request.UserId, request.Vertical.String(), common.UserHistory{
			ApMin:       2,
			ApMax:       7,
			Sensitivity: 4,
			UpdatedTs:   timeHelpers.FormatTime(time.Now(), timeHelpers.DateInYYYYMMDDFormat),
		})
	}
	// 2. Save the current realtime data
	realTimeData := common.RealTimeData{
		Org:         request.GetOrg(),
		Vertical:    request.GetVertical(),
		PageHitType: request.GetPageType(),
		UserId:      request.GetUserId(),
		EntityId:    request.GetEntityId(),
	}
	realTimeData.RequestDate, _ = strconv.Atoi(timeHelpers.GetRequiredTimeFormat(request.GetRequestDate(), timeHelpers.DateTimeFormat, timeHelpers.DateInYYYYMMDDFormat))
	realTimeData.TravelDate, _ = strconv.Atoi(timeHelpers.GetRequiredTimeFormat(request.GetTravelStartDate(), timeHelpers.DateInYYYYMMDDFormat, timeHelpers.DateInYYYYMMDDFormat))
	travelStartDate, _ := strconv.Atoi(timeHelpers.GetRequiredTimeFormat(request.GetTravelStartDate(), timeHelpers.DateInYYYYMMDDFormat, timeHelpers.DateInYYYYMMDDFormat))
	travelEndDate, _ := strconv.Atoi(timeHelpers.GetRequiredTimeFormat(request.GetTravelEndDate(), timeHelpers.DateInYYYYMMDDFormat, timeHelpers.DateInYYYYMMDDFormat))
	realTimeData.RoomNights = travelEndDate - travelStartDate
	userData, err := repo.SaveRealTimeData(ctx, realTimeData)
	log.Print(userData)

	// 3. Calculate the score by considering the UserData --> RealTimeDataResponse, entityData, UserHistory

	// 4. Push the score to manthan

	return err
}
