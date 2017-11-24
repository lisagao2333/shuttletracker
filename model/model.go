// Package model provides structs used by multiple packages within shuttletracker.
package model

import (
	"time"
)

// VehicleUpdate represents a single position observed for a Vehicle.
type Update struct {
	ID        int64
	VehicleID int64     `json:"vehicleID,string"   bson:"vehicleID,omitempty" db:"vehicle_id"`
	Latitude  float64   `json:"lat"         bson:"lat"`
	Longitude float64   `json:"lng"         bson:"lng"`
	Heading   string    `json:"heading"     bson:"heading"`
	Speed     string    `json:"speed"       bson:"speed"`
	Lock      string    `json:"lock"        bson:"lock"`
	Timestamp time.Time `json:"time"        bson:"time"`
	Created   time.Time `json:"created"     bson:"created"`
	RouteID   int64     `json:"RouteID,string"     bson:"routeID"`
}

// Vehicle represents an object being tracked.
type Vehicle struct {
	ID      int64     `json:"vehicleID,string"`
	ITrakID int64     `json:"itrakID"   bson:"itrakID,omitempty" db:"itrak_id"`
	Name    string    `json:"vehicleName" bson:"vehicleName"`
	Created time.Time `bson:"created"`
	Updated time.Time `bson:"updated"`
	Enabled bool      `json:"enabled"     bson:"enabled"`
}

// Status contains a detailed message on the tracked object's status.
type Status struct {
	Public  bool      `bson:"public"`
	Message string    `json:"message" bson:"message"`
	Created time.Time `bson:"created"`
	Updated time.Time `bson:"updated"`
}

type LatestPosition struct {
	Longitude     string    `json:"longitude"`
	Latitude      string    `json:"latitude"`
	Timestamp     time.Time `json:"timestamp"`
	Speed         float64   `json:"speed"`
	Heading       int       `json:"heading"`
	Cardinal      string    `json:"cardinal_point"`
	StatusMessage *string   `json:"public_status_message"` // this is a pointer so it defaults to null
}

// Coord represents a single lat/lng point used to draw routes
type Coord struct {
	Lat float64 `json:"lat" bson:"lat"`
	Lng float64 `json:"lng" bson:"lng"`
}

// Route represents a set of coordinates to draw a path on our tracking map
type Route struct {
	ID          int64       `json:"id,string"             bson:"id" db:"id"`
	Name        string      `json:"name"           bson:"name" db:"name"`
	Description string      `json:"description"    bson:"description"`
	StartTime   string      `json:"startTime"      bson:"startTime"`
	EndTime     string      `json:"endTime"        bson:"endTime"`
	Enabled     bool        `json:"enabled,bool"	  bson:"enabled"`
	Color       string      `json:"color"          bson:"color"`
	Width       int         `json:"width,string"   bson:"width"`
	Coords      []Coord     `json:"coords"         bson:"coords"`
	Duration    []Segment   `json:"duration"       bson:"duration"`
	StopsID     []int64     `json:"stopsid,[]string"        bson:"stopsid"`
	Created     time.Time   `json:"created"        bson:"created" db:"created"`
	Updated     time.Time   `json:"updated"        bson:"updated" db:"updated"`
	Stops       []RouteStop `json:"stops"`
}

// Stop indicates where a tracked object is scheduled to arrive
type Stop struct {
	ID           int64   `json:"id,string"             bson:"id"`
	Name         string  `json:"name"           bson:"name"`
	Description  string  `json:"description"    bson:"description"`
	Latitude     float64 `json:"lat,string"     bson:"lat"`
	Longitude    float64 `json:"lng,string"     bson:"lng"`
	Address      string  `json:"address"        bson:"address"`
	StartTime    string  `json:"startTime"      bson:"startTime"`
	EndTime      string  `json:"endTime"        bson:"endTime"`
	Enabled      bool    `json:"enabled,string" bson:"enabled"`
	SegmentIndex int     `json:"segmentindex"   bson:"segmentindex"`
	Order        int64   `db:"stop_order"`
	Created      time.Time
	Updated      time.Time
	Routes       []int64 `json:"routes" db:"routes"`
}

type RouteStop struct {
	ID      int64 `json:"id"`
	RouteID int64 `json:"routeID"`
	StopID  int64 `json:"stopID"`
	Order   int64 `json:"order" db:"stop_order"`
}

type MapPoint struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
type MapResponsePoint struct {
	Location      MapPoint `json:"location"`
	OriginalIndex int      `json:"originalIndex,omitempty"`
	PlaceID       string   `json:"placeId"`
}
type MapResponse struct {
	SnappedPoints []MapResponsePoint
}

type MapDistanceMatrixDuration struct {
	Value int    `json:"value"`
	Text  string `json:"text"`
}

type MapDistanceMatrixDistance struct {
	Value int    `json:"value"`
	Text  string `json:"text"`
}

type MapDistanceMatrixElement struct {
	Status   string                    `json:"status"`
	Duration MapDistanceMatrixDuration `json:"duration"`
	Distance MapDistanceMatrixDistance `json:"distance"`
}

type MapDistanceMatrixElements struct {
	Elements []MapDistanceMatrixElement `json:"elements"`
}
type MapDistanceMatrixResponse struct {
	Status               string                      `json:"status"`
	OriginAddresses      []string                    `json:"origin_addresses"`
	DestinationAddresses []string                    `json:"destination_addresses"`
	Rows                 []MapDistanceMatrixElements `json:"rows"`
}

type Segment struct {
	ID       string   `json:"id"`
	Start    MapPoint `json:"origin"`
	End      MapPoint `json:"destination"`
	Distance float64  `json:"distance"`
	Duration float64  `json:"duration"`
}
