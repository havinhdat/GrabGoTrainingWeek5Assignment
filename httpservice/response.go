package httpservice

import (
	pwc "grab/week5/GrabGoTrainingWeek5Assignment/postwithcomment"
)

type PostWithCommentsResponse struct {
	Posts []pwc.PostWithComment `json:"posts"`
}
