package entity

import (
	"github.com/adiatma85/own-go-sdk/mongo"
	"github.com/adiatma85/own-go-sdk/null"
	"go.mongodb.org/mongo-driver/bson"
)

const AggridCollectionName = "aggrid"

type AgGridMetricParam struct {
	ID             string  `json:"id" bson:"_id,omitempty"`
	FarmID         int64   `json:"farmID" bson:"farm_id,omitempty"`
	PondID         int64   `json:"pondID" bson:"pond_id,omitempty"`
	CycleID        int64   `json:"cycleID" bson:"cycle_id,omitempty"`
	MetricDate     bson.M  `bson:"metric_date,omitempty" json:"metricDate"`
	PondName       string  `bson:"pond_name,omitempty" json:"pondName"`
	StockingActual float64 `bson:"stocking_actual,omitempty" json:"stockingActual"`
	Doc            int64   `bson:"doc,omitempty" json:"doc"`
	Dop            int64   `bson:"dop,omitempty" json:"dop"`
	TemplateType   string  `bson:"template_type,omitempty" json:"templateType"`

	// Secondary
	YAxis                null.Int64 `json:"y_axis" bson:"y_axis,omitempty"` // notate what is row number
	BypassCache          bool       `db:"-" json:"-" bson:"-"`
	PaginationParamMongo `bson:"-"`
	QueryOption          mongo.Option `bson:"-"`
}

type AgGridMetric struct {
	ID           string      `json:"id" bson:"_id,omitempty"`
	FarmID       null.Int64  `json:"farmID" bson:"farm_id,omitempty"`
	CycleID      null.Int64  `json:"cycleID" bson:"cycle_id,omitempty"`
	PondID       null.Int64  `json:"pondID" bson:"pond_id,omitempty"`
	TemplateType null.String `json:"templateType" bson:"template_type,omitempty"`

	// Y Axis
	YAxis null.Int64 `name:"y_axis" json:"y_axis" bson:"y_axis,omitempty"`

	// General Attributes
	MetricDate     null.Date    `name:"metric_date" bson:"metric_date,omitempty" json:"metricDate"`
	PondName       null.String  `name:"pond_name" bson:"pond_name,omitempty" json:"pondName"`
	StockingActual null.Float64 `name:"stocking_actual" bson:"stocking_actual,omitempty" json:"stockingActual"`
	Doc            null.Int64   `name:"doc" bson:"doc,omitempty" json:"doc"`
	Dop            null.Int64   `name:"dop" bson:"dop,omitempty" json:"dop"`

	// Farm Technician Parameters
	WaterHeightMorning             null.Float64 `name:"water_height_morning" bson:"water_height_morning,omitempty" json:"waterHeightMorning"`
	WaterHeight                    null.Float64 `name:"water_height" bson:"water_height,omitempty" json:"waterHeight"`
	WaterColorMorning              null.String  `name:"water_color_morning" bson:"water_color_morning,omitempty" json:"waterColorMorning"`
	WaterColorEvening              null.String  `name:"water_color_evening" bson:"water_color_evening,omitempty" json:"waterColorEvening"`
	BrightnessMorning              null.Float64 `name:"brightness_morning" bson:"brightness_morning,omitempty" json:"brightnessMorning"`
	BrightnessEvening              null.Float64 `name:"brightness_evening" bson:"brightness_evening,omitempty" json:"brightnessEvening"`
	WeatherMorning                 null.String  `name:"weather_morning" bson:"weather_morning,omitempty" json:"weatherMorning"`
	WeatherEvening                 null.String  `name:"weather_evening" bson:"weather_evening,omitempty" json:"weatherEvening"`
	WeatherNight                   null.String  `name:"weather_night" bson:"weather_night,omitempty" json:"weatherNight"`
	SyphonHourMorning              null.Float64 `name:"syphon_hour_morning" bson:"syphon_hour_morning,omitempty" json:"syphonHourMorning"`
	SyphonHourEvening              null.Float64 `name:"syphon_hour_evening" bson:"syphon_hour_evening,omitempty" json:"syphonHourEvening"`
	WaterOut                       null.Float64 `name:"water_out" bson:"water_out,omitempty" json:"waterOut"`
	WaterIn                        null.Float64 `name:"water_in" bson:"water_in,omitempty" json:"waterIn"`
	WaterOverflowDuration          null.Float64 `name:"water_overflow_duration" bson:"water_overflow_duration,omitempty" json:"waterOverflowDuration"`
	MortalityDaily                 null.Float64 `name:"mortality_daily" bson:"mortality_daily,omitempty" json:"mortalityDaily"`
	MortalityCum                   null.Float64 `name:"mortality_cum" bson:"mortality_cum,omitempty" json:"mortalityCum"`
	PartialHarvest                 null.Float64 `name:"partial_harvest" bson:"partial_harvest,omitempty" json:"partialHarvest"`
	SrTodayPercentage              null.Float64 `name:"sr_today_percentage" bson:"sr_today_percentage,omitempty" json:"srTodayPercentage"`
	WaterWheel1HpMorning           null.Float64 `name:"water_wheel_1hp_morning" bson:"water_wheel_1hp_morning,omitempty" json:"waterWheel1hpMorning"`
	WaterWheel1HpMorningDuration   null.Float64 `name:"water_wheel_1hp_morning_duration" bson:"water_wheel_1hp_morning_duration,omitempty" json:"waterWheel1hpMorningDuration"`
	WaterWheel1HpAfternoon         null.Float64 `name:"water_wheel_1hp_afternoon" bson:"water_wheel_1hp_afternoon,omitempty" json:"waterWheel1hpAfternoon"`
	WaterWheel1HpAfternoonDuration null.Float64 `name:"water_wheel_1hp_afternoon_duration" bson:"water_wheel_1hp_afternoon_duration,omitempty" json:"waterWheel1hpAfternoonDuration"`
	WaterWheel2HpMorning           null.Float64 `name:"water_wheel_2hp_morning" bson:"water_wheel_2hp_morning,omitempty" json:"waterWheel2hpMorning"`
	WaterWheel2HpMorningDuration   null.Float64 `name:"water_wheel_2hp_morning_duration" bson:"water_wheel_2hp_morning_duration,omitempty" json:"waterWheel2hpMorningDuration"`
	WaterWheel2HpAfternoon         null.Float64 `name:"water_wheel_2hp_afternoon" bson:"water_wheel_2hp_afternoon,omitempty" json:"waterWheel2hpAfternoon"`
	WaterWheel2HpAfternoonDuration null.Float64 `name:"water_wheel_2hp_afternoon_duration" bson:"water_wheel_2hp_afternoon_duration,omitempty" json:"waterWheel2hpAfternoonDuration"`
	BlowerMorning                  null.Float64 `name:"blower_morning" bson:"blower_morning,omitempty" json:"blowerMorning"`
	BlowerMorningDuration          null.Float64 `name:"blower_morning_duration" bson:"blower_morning_duration,omitempty" json:"blowerMorningDuration"`
	BlowerEvening                  null.Float64 `name:"blower_evening" bson:"blower_evening,omitempty" json:"blowerAfternoon"`
	BlowerAfternoongDuration       null.Float64 `name:"blower_afternoon_duration" bson:"blower_afternoon_duration,omitempty" json:"blowerAfternoonDuration"`
	TotalPowerMorning              null.Float64 `name:"total_power_morning" bson:"total_power_morning,omitempty" json:"totalPowerMorning"`
	TotalPowerEvening              null.Float64 `name:"total_power_evening" bson:"total_power_evening,omitempty" json:"totalPowerAfternoon"`
	BrandFeed                      null.String  `name:"brand_feed" bson:"brand_feed,omitempty" json:"brandFeed"`
	FeedCode                       null.String  `name:"feed_code" bson:"feed_code,omitempty" json:"feedCode"`
	FeedSession1Weight             null.Float64 `name:"feed_session1_weight" bson:"feed_session1_weight,omitempty" json:"feedSession1Weight"`
	FeedSession2Weight             null.Float64 `name:"feed_session2_weight" bson:"feed_session2_weight,omitempty" json:"feedSession2Weight"`
	FeedSession3Weight             null.Float64 `name:"feed_session3_weight" bson:"feed_session3_weight,omitempty" json:"feedSession3Weight"`
	FeedSession4Weight             null.Float64 `name:"feed_session4_weight" bson:"feed_session4_weight,omitempty" json:"feedSession4Weight"`
	FeedSession5Weight             null.Float64 `name:"feed_session5_weight" bson:"feed_session5_weight,omitempty" json:"feedSession5Weight"`
	FeedSession6Weight             null.Float64 `name:"feed_session6_weight" bson:"feed_session6_weight,omitempty" json:"feedSession6Weight"`
	BrandFeedMix                   null.String  `name:"brand_feed_mix" bson:"brand_feed_mix,omitempty" json:"brandFeedMix"`
	FeedCodeMix                    null.String  `name:"brand_feed_code_mix" bson:"brand_feed_code_mix,omitempty" json:"feedCodeMix"`
	FeedSession1WeightMix          null.Float64 `name:"feed_session1_weight_mix" bson:"feed_session1_weight_mix,omitempty" json:"feedSession1WeightMix"`
	FeedSession2WeightMix          null.Float64 `name:"feed_session2_weight_mix" bson:"feed_session2_weight_mix,omitempty" json:"feedSession2WeightMix"`
	FeedSession3WeightMix          null.Float64 `name:"feed_session3_weight_mix" bson:"feed_session3_weight_mix,omitempty" json:"feedSession3WeightMix"`
	FeedSession4WeightMix          null.Float64 `name:"feed_session4_weight_mix" bson:"feed_session4_weight_mix,omitempty" json:"feedSession4WeightMix"`
	FeedSession5WeightMix          null.Float64 `name:"feed_session5_weight_mix" bson:"feed_session5_weight_mix,omitempty" json:"feedSession5WeightMix"`
	FeedSession6WeightMix          null.Float64 `name:"feed_session6_weight_mix" bson:"feed_session6_weight_mix,omitempty" json:"feedSession6WeightMix"`
	TotalFeedActual                null.Float64 `name:"total_feed_actual" bson:"total_feed_actual,omitempty" json:"totalFeedActual"`
	TotalCumFeed                   null.Float64 `name:"total_cum_feed" bson:"total_cum_feed,omitempty" json:"totalCumFeed"`
	FeedNote                       null.String  `name:"feed_note_session1" bson:"feed_note_session1,omitempty" json:"feedNote"`
	ShrimpCondition                null.String  `name:"shrimp_condition_session1" bson:"shrimp_condition_session1,omitempty" json:"shrimpCondition"`
	AncoPercentage                 null.Float64 `name:"anco_percentage" bson:"anco_percentage,omitempty" json:"ancoPercentage"`
	AncoSess1Anco1                 null.String  `name:"anco_session1_anco1" bson:"anco_session1_anco1,omitempty" json:"ancoSession1Anco1"`
	AncoSess1Anco2                 null.String  `name:"anco_session1_anco2" bson:"anco_session1_anco2,omitempty" json:"ancoSession1Anco2"`
	AncoSess1Anco3                 null.String  `name:"anco_session1_anco3" bson:"anco_session1_anco3,omitempty" json:"ancoSession1Anco3"`
	AncoSess1Anco4                 null.String  `name:"anco_session1_anco4" bson:"anco_session1_anco4,omitempty" json:"ancoSession1Anco4"`
	AncoSess1Anco5                 null.String  `name:"anco_session1_anco5" bson:"anco_session1_anco5,omitempty" json:"ancoSession1Anco5"`
	AncoSess2Anco1                 null.String  `name:"anco_session2_anco1" bson:"anco_session2_anco1,omitempty" json:"ancoSession2Anco1"`
	AncoSess2Anco2                 null.String  `name:"anco_session2_anco2" bson:"anco_session2_anco2,omitempty" json:"ancoSession2Anco2"`
	AncoSess2Anco3                 null.String  `name:"anco_session2_anco3" bson:"anco_session2_anco3,omitempty" json:"ancoSession2Anco3"`
	AncoSess2Anco4                 null.String  `name:"anco_session2_anco4" bson:"anco_session2_anco4,omitempty" json:"ancoSession2Anco4"`
	AncoSess2Anco5                 null.String  `name:"anco_session2_anco5" bson:"anco_session2_anco5,omitempty" json:"ancoSession2Anco5"`
	AncoSess3Anco1                 null.String  `name:"anco_session3_anco1" bson:"anco_session3_anco1,omitempty" json:"ancoSession3Anco1"`
	AncoSess3Anco2                 null.String  `name:"anco_session3_anco2" bson:"anco_session3_anco2,omitempty" json:"ancoSession3Anco2"`
	AncoSess3Anco3                 null.String  `name:"anco_session3_anco3" bson:"anco_session3_anco3,omitempty" json:"ancoSession3Anco3"`
	AncoSess3Anco4                 null.String  `name:"anco_session3_anco4" bson:"anco_session3_anco4,omitempty" json:"ancoSession3Anco4"`
	AncoSess3Anco5                 null.String  `name:"anco_session3_anco5" bson:"anco_session3_anco5,omitempty" json:"ancoSession3Anco5"`
	Treatment1Brand                null.String  `name:"treatment1_brand" bson:"treatment_1_brand,omitempty" json:"treatment1Brand"`
	Treatment1Type                 null.String  `name:"treatment1_type" bson:"treatment_1_type,omitempty" json:"treatment1Type"`
	Treatment1Total                null.Float64 `name:"treatment1_total" bson:"treatment_1_total,omitempty" json:"treatment1Total"`
	Treatment1Unit                 null.String  `name:"treatment1_unit" bson:"treatment_1_unit,omitempty" json:"treatment1Unit"`
	Treatment2Brand                null.String  `name:"treatment2_brand" bson:"treatment_2_brand,omitempty" json:"treatment2Brand"`
	Treatment2Type                 null.String  `name:"treatment2_type" bson:"treatment_2_type,omitempty" json:"treatment2Type"`
	Treatment2Total                null.Float64 `name:"treatment2_total" bson:"treatment_2_total,omitempty" json:"treatment2Total"`
	Treatment2Unit                 null.String  `name:"treatment2_unit" bson:"treatment_2_unit,omitempty" json:"treatment2Unit"`
	Treatment3Brand                null.String  `name:"treatment3_brand" bson:"treatment_3_brand,omitempty" json:"treatment3Brand"`
	Treatment3Type                 null.String  `name:"treatment3_type" bson:"treatment_3_type,omitempty" json:"treatment3Type"`
	Treatment3Total                null.Float64 `name:"treatment3_total" bson:"treatment_3_total,omitempty" json:"treatment3Total"`
	Treatment3Unit                 null.String  `name:"treatment3_unit" bson:"treatment_3_unit,omitempty" json:"treatment3Unit"`
	Treatment4Brand                null.String  `name:"treatment4_brand" bson:"treatment_4_brand,omitempty" json:"treatment4Brand"`
	Treatment4Type                 null.String  `name:"treatment4_type" bson:"treatment_4_type,omitempty" json:"treatment4Type"`
	Treatment4Total                null.Float64 `name:"treatment4_total" bson:"treatment_4_total,omitempty" json:"treatment4Total"`
	Treatment4Unit                 null.String  `name:"treatment4_unit" bson:"treatment_4_unit,omitempty" json:"treatment4Unit"`
	Treatment5Brand                null.String  `name:"treatment5_brand" bson:"treatment_5_brand,omitempty" json:"treatment5Brand"`
	Treatment5Type                 null.String  `name:"treatment5_type" bson:"treatment_5_type,omitempty" json:"treatment5Type"`
	Treatment5Total                null.Float64 `name:"treatment5_total" bson:"treatment_5_total,omitempty" json:"treatment5Total"`
	Treatment5Unit                 null.String  `name:"treatment5_unit" bson:"treatment_5_unit,omitempty" json:"treatment5Unit"`
	Treatment6Brand                null.String  `name:"treatment6_brand" bson:"treatment_6_brand,omitempty" json:"treatment6Brand"`
	Treatment6Type                 null.String  `name:"treatment6_type" bson:"treatment_6_type,omitempty" json:"treatment6Type"`
	Treatment6Total                null.Float64 `name:"treatment6_total" bson:"treatment_6_total,omitempty" json:"treatment6Total"`
	Treatment6Unit                 null.String  `name:"treatment6_unit" bson:"treatment_6_unit,omitempty" json:"treatment6Unit"`
	Treatment7Brand                null.String  `name:"treatment7_brand" bson:"treatment_7_brand,omitempty" json:"treatment7Brand"`
	Treatment7Type                 null.String  `name:"treatment7_type" bson:"treatment_7_type,omitempty" json:"treatment7Type"`
	Treatment7Total                null.Float64 `name:"treatment7_total" bson:"treatment_7_total,omitempty" json:"treatment7Total"`
	Treatment7Unit                 null.String  `name:"treatment7_unit" bson:"treatment_7_unit,omitempty" json:"treatment7Unit"`
	Treatment8Brand                null.String  `name:"treatment8_brand" bson:"treatment_8_brand,omitempty" json:"treatment8Brand"`
	Treatment8Type                 null.String  `name:"treatment8_type" bson:"treatment_8_type,omitempty" json:"treatment8Type"`
	Treatment8Total                null.Float64 `name:"treatment8_total" bson:"treatment_8_total,omitempty" json:"treatment8Total"`
	Treatment8Unit                 null.String  `name:"treatment8_unit" bson:"treatment_8_unit,omitempty" json:"treatment8Unit"`
	Treatment9Brand                null.String  `name:"treatment9_brand" bson:"treatment_9_brand,omitempty" json:"treatment9Brand"`
	Treatment9Type                 null.String  `name:"treatment9_type" bson:"treatment_9_type,omitempty" json:"treatment9Type"`
	Treatment9Total                null.Float64 `name:"treatment9_total" bson:"treatment_9_total,omitempty" json:"treatment9Total"`
	Treatment9Unit                 null.String  `name:"treatment9_unit" bson:"treatment_9_unit,omitempty" json:"treatment9Unit"`
	Treatment10Brand               null.String  `name:"treatment10_brand" bson:"treatment_10_brand,omitempty" json:"treatment10Brand"`
	Treatment10Type                null.String  `name:"treatment10_type" bson:"treatment_10_type,omitempty" json:"treatment10Type"`
	Treatment10Total               null.Float64 `name:"treatment10_total" bson:"treatment_10_total,omitempty" json:"treatment10Total"`
	Treatment10Unit                null.String  `name:"treatment10_unit" bson:"treatment_10_unit,omitempty" json:"treatment10Unit"`
	TreatmentNote                  null.String  `name:"treatment_note" bson:"treatment_note,omitempty" json:"treatmentNote"`

	BrandFeedSession2 null.String `name:"brand_feed_session2" bson:"brand_feed_session_2,omitempty" json:"brandFeedSession2"`
	FeedCodeSession2  null.String `name:"feed_code_session2" bson:"feed_code_session_2,omitempty" json:"feedCodeSession2"`
	BrandFeedSession3 null.String `name:"brand_feed_session3" bson:"brand_feed_session_3,omitempty" json:"brandFeedSession3"`
	FeedCodeSession3  null.String `name:"feed_code_session3" bson:"feed_code_session_3,omitempty" json:"feedCodeSession3"`
	BrandFeedSession4 null.String `name:"brand_feed_session4" bson:"brand_feed_session_4,omitempty" json:"brandFeedSession4"`
	FeedCodeSession4  null.String `name:"feed_code_session4" bson:"feed_code_session_4,omitempty" json:"feedCodeSession4"`
	BrandFeedSession5 null.String `name:"brand_feed_session5" bson:"brand_feed_session_5,omitempty" json:"brandFeedSession5"`
	FeedCodeSession5  null.String `name:"feed_code_session5" bson:"feed_code_session_5,omitempty" json:"feedCodeSession5"`
	BrandFeedSession6 null.String `name:"brand_feed_session6" bson:"brand_feed_session_6,omitempty" json:"brandFeedSession6"`
	FeedCodeSession6  null.String `name:"feed_code_session6" bson:"feed_code_session_6,omitempty" json:"feedCodeSession6"`

	BrandFeedMixSession2     null.String `name:"brand_feed_mix_session2" bson:"brand_feed_mix_session_2,omitempty" json:"brandFeedMixSession2"`
	BrandFeedCodeMixSession2 null.String `name:"brand_feed_code_mix_session2" bson:"brand_feed_code_mix_session_2,omitempty" json:"brandFeedCodeMixSession2"`
	BrandFeedMixSession3     null.String `name:"brand_feed_mix_session3" bson:"brand_feed_mix_session_3,omitempty" json:"brandFeedMixSession3"`
	BrandFeedCodeMixSession3 null.String `name:"brand_feed_code_mix_session3" bson:"brand_feed_code_mix_session_3,omitempty" json:"brandFeedCodeMixSession3"`
	BrandFeedMixSession4     null.String `name:"brand_feed_mix_session4" bson:"brand_feed_mix_session_4,omitempty" json:"brandFeedMixSession4"`
	BrandFeedCodeMixSession4 null.String `name:"brand_feed_code_mix_session4" bson:"brand_feed_code_mix_session_4,omitempty" json:"brandFeedCodeMixSession4"`
	BrandFeedMixSession5     null.String `name:"brand_feed_mix_session5" bson:"brand_feed_mix_session_5,omitempty" json:"brandFeedMixSession5"`
	BrandFeedCodeMixSession5 null.String `name:"brand_feed_code_mix_session5" bson:"brand_feed_code_mix_session_5,omitempty" json:"brandFeedCodeMixSession5"`
	BrandFeedMixSession6     null.String `name:"brand_feed_mix_session6" bson:"brand_feed_mix_session_6,omitempty" json:"brandFeedMixSession6"`
	BrandFeedCodeMixSession6 null.String `name:"brand_feed_code_mix_session6" bson:"brand_feed_code_mix_session_6,omitempty" json:"brandFeedCodeMixSession6"`

	// Farm Technician Parameter End

	// Lab analyst Parameters
	// "Lab Analyst" sheet

	// SAL is already on Farm Technician
	Sal                          null.Float64 `name:"sal" bson:"sal,omitempty" json:"sal"`
	SalTime                      null.String  `name:"salinity_time" bson:"salinity_time,omitempty" json:"salTime"`
	DOMorning                    null.Float64 `name:"do_time0500" bson:"do_time0500,omitempty" json:"doTime0500"`
	DOMorningTime                null.String  `name:"do_time0500_time" bson:"do_time0500_time,omitempty" json:"doTime0500Time"`
	DOAfternoon                  null.Float64 `name:"do_time1300" bson:"do_time1300,omitempty" json:"doTime1300"`
	DOAfternoonTime              null.String  `name:"do_time1300_time" bson:"do_time1300_time,omitempty" json:"doTime1300Time"`
	DONight                      null.Float64 `name:"do_time2100" bson:"do_time2100,omitempty" json:"doTime2100"`
	DONightTime                  null.String  `name:"do_time2100_time" bson:"do_time2100_time,omitempty" json:"doTime2100Time"`
	PHMorning                    null.Float64 `name:"ph_time0430" bson:"ph_time0430,omitempty" json:"phTime0430"`
	PHMorningTime                null.String  `name:"ph_time0430_time" bson:"ph_time0430_time,omitempty" json:"phTime0430Time"`
	PHAfternoon                  null.Float64 `name:"ph_time1300" bson:"ph_time1300,omitempty" json:"phTime1300"`
	PHAfternoonTime              null.String  `name:"ph_time1300_time" bson:"ph_time1300_time,omitempty" json:"phTime1300Time"`
	RangePH                      null.Float64 `name:"range_ph" bson:"range_ph,omitempty" json:"rangePH"`
	TemperatureMorning           null.Float64 `name:"temperature_time0500" bson:"temperature_time0500,omitempty" json:"temperatureTime0500"`
	TemperatureMorningTime       null.String  `name:"temperature_time0500_time" bson:"temperature_time0500_time,omitempty" json:"temperatureTime0500Time"`
	TemperatureAfternoon         null.Float64 `name:"temperature_time1400" bson:"temperature_time1400,omitempty" json:"temperatureTime1400"`
	TemperatureAfternoonTime     null.String  `name:"temperature_time1400_time" bson:"temperature_time1400_time,omitempty" json:"temperatureTime1400Time"`
	TemperatureNight             null.Float64 `name:"temperature_time2100" bson:"temperature_time2100,omitempty" json:"temperatureTime2100"`
	TemperatureNightTime         null.String  `name:"temperature_time2100_time" bson:"temperature_time2100_time,omitempty" json:"temperatureTime2100Time"`
	ORPMorning                   null.Float64 `name:"orp_time0500" bson:"orp_time0500,omitempty" json:"orpTime0500"`
	Tan                          null.Float64 `name:"tan" bson:"tan,omitempty" json:"tan"`
	Nh3                          null.Float64 `name:"nh3" bson:"nh3,omitempty" json:"nh3"`
	Nh4                          null.Float64 `name:"nh4" bson:"nh4,omitempty" json:"nh4"`
	No3                          null.Float64 `name:"no3" bson:"no3,omitempty" json:"no3"`
	No2                          null.Float64 `name:"no2" bson:"no2,omitempty" json:"no2"`
	Po4                          null.Float64 `name:"po4" bson:"po4,omitempty" json:"po4"`
	RatioNP                      null.Float64 `name:"ratio_np" bson:"ratio_np,omitempty" json:"ratioNP"`
	Oh                           null.Float64 `name:"oh" bson:"oh,omitempty" json:"oh"`
	Co3                          null.Float64 `name:"co3" bson:"co3,omitempty" json:"co3"`
	Hco3                         null.Float64 `name:"hco3" bson:"hco3,omitempty" json:"hco3"`
	TAlk                         null.Float64 `name:"t_alk" bson:"t_alk,omitempty" json:"tAlk"`
	Ca                           null.Float64 `name:"ca" bson:"ca,omitempty" json:"ca"`
	Mg                           null.Float64 `name:"mg" bson:"mg,omitempty" json:"mg"`
	MgCa                         null.Float64 `name:"mgca" bson:"mgca,omitempty" json:"mgca"`
	Camg                         null.Float64 `name:"camg" bson:"camg,omitempty" json:"camg"`
	THard                        null.Float64 `name:"t_hard" bson:"t_hard,omitempty" json:"tHard"`
	Tom                          null.Float64 `name:"tom" bson:"tom,omitempty" json:"tom"`
	BacteryGaPercentage          null.Float64 `name:"bactery_ga_percentage" bson:"bactery_ga_percentage,omitempty" json:"bacteryGaPercentage"`
	BacteryBgaPercentage         null.Float64 `name:"bactery_bga_percentage" bson:"bactery_bga_percentage,omitempty" json:"bacteryBgaPercentage"`
	BacteryDiaPercentage         null.Float64 `name:"bactery_dia_percentage" bson:"bactery_dia_percentage,omitempty" json:"bacteryDiaPercentage"`
	BacteryEuglenoPercentage     null.Float64 `name:"bactery_eugleno_percentage" bson:"bactery_eugleno_percentage,omitempty" json:"bacteryEuglenoPercentage"`
	BacteryDinoPercentage        null.Float64 `name:"bactery_dino_percentage" bson:"bactery_dino_percentage,omitempty" json:"bacteryDinoPercentage"`
	BacteryZooPercentage         null.Float64 `name:"bactery_zoo_percentage" bson:"bactery_zoo_percentage,omitempty" json:"bacteryZooPercentage"`
	TotalOtherPlanktonPercentage null.Float64 `name:"total_other_plankton_percentage" bson:"total_other_plankton_percentage,omitempty" json:"totalOtherPlanktonPercentage"`
	BacteryVibrioY               null.Float64 `name:"bactery_vibrio_y" bson:"bactery_vibrio_y,omitempty" json:"bacteryVibrioY"`
	BacteryVibrioG               null.Float64 `name:"bactery_vibrio_g" bson:"bactery_vibrio_g,omitempty" json:"bacteryVibrioG"`
	BacteryBlack                 null.Float64 `name:"bactery_black" bson:"bactery_black,omitempty" json:"bacteryBlack"`
	BacteryVibrioTvc             null.Float64 `name:"bactery_vibrio_tvc" bson:"bactery_vibrio_tvc,omitempty" json:"bacteryVibrioTvc"`
	BacteryTbc                   null.Float64 `name:"bactery_tbc" bson:"bactery_tbc,omitempty" json:"bacteryTbc"`
	TvcTbcPercentage             null.Float64 `name:"tvc_tbc_percentage" bson:"tvc_tbc_percentage,omitempty" json:"tvcTbcPercentage"`
	FunctionalNp                 null.Float64 `name:"functional_np" bson:"functional_np,omitempty" json:"functionalNp"`

	// "Lab Calculation Sheet"
	H2so4Ap                null.Float64 `name:"h2so4_ap" bson:"h2so4_ap,omitempty" json:"h2so4Ap"`
	H2so4Bt                null.Float64 `name:"h2so4_bt" bson:"h2so4_bt,omitempty" json:"h2so4Bt"`
	NAsam                  null.Float64 `name:"n_asam" bson:"n_asam,omitempty" json:"nAsam"`
	SampleVolumeAlkalinity null.Float64 `name:"sample_volume_alkalinity" bson:"sample_volume_alkalinity,omitempty" json:"sampleVolumeAlkalinity"`
	PtAlkalinityDiff       null.Float64 `name:"pt_alkalinity_diff" bson:"pt_alkalinity_diff,omitempty" json:"ptAlkalinityDiff"`
	PAlkalinity            null.Float64 `name:"p_alkalinity" bson:"p_alkalinity,omitempty" json:"pAlkalinity"`
	TAlkalinity            null.Float64 `name:"t_alkalinity" bson:"t_alkalinity,omitempty" json:"tAlkalinity"`
	NitritDilutionFactor   null.Float64 `name:"nitrit_dilution_factor" bson:"nitrit_dilution_factor,omitempty" json:"nitritDilutionFactor"`
	NitritBlanko           null.Float64 `name:"nitrit_blanko" bson:"nitrit_blanko,omitempty" json:"nitritBlanko"`
	NitritAbsAs            null.Float64 `name:"nitrit_abs_as" bson:"nitrit_abs_as,omitempty" json:"nitritAbsAs"`
	NitritSampleVolume     null.Float64 `name:"nitrit_sample_volume" bson:"nitrit_sample_volume,omitempty" json:"nitritSampleVolume"`
	NitritConcC            null.Float64 `name:"nitrit_conc_c" bson:"nitrit_conc_c,omitempty" json:"nitritConcC"`
	NitritDescription      null.Float64 `name:"nitrit_description" bson:"nitrit_description,omitempty" json:"nitritDescription"`
	FosfatDilutionFactor   null.Float64 `name:"fosfat_dilution_factor" bson:"fosfat_dilution_factor,omitempty" json:"fosfatDilutionFactor"`
	FosfatBlanko           null.Float64 `name:"fosfat_blanko" bson:"fosfat_blanko,omitempty" json:"fosfatBlanko"`
	FosfatAbsAs            null.Float64 `name:"fosfat_abs_as" bson:"fosfat_abs_as,omitempty" json:"fosfatAbsAs"`
	FosfatSampleVolume     null.Float64 `name:"fosfat_sample_volume" bson:"fosfat_sample_volume,omitempty" json:"fosfatSampleVolume"`
	FosfatConcC            null.Float64 `name:"fosfat_conc_c" bson:"fosfat_conc_c,omitempty" json:"fosfatConcC"`
	FosfatDescription      null.Float64 `name:"fosfat_description" bson:"fosfat_description,omitempty" json:"fosfatDescription"`
	TanDilutionFactor      null.Float64 `name:"tan_dilution_factor" bson:"tan_dilution_factor,omitempty" json:"tanDilutionFactor"`
	TanBlanko              null.Float64 `name:"tan_blanko" bson:"tan_blanko,omitempty" json:"tanBlanko"`
	TanAbsAs               null.Float64 `name:"tan_abs_as" bson:"tan_abs_as,omitempty" json:"tanAbsAs"`
	TanSampleVolume        null.Float64 `name:"tan_sample_volume" bson:"tan_sample_volume,omitempty" json:"tanSampleVolume"`
	TanConcC               null.Float64 `name:"tan_conc_c" bson:"tan_conc_c,omitempty" json:"tanConcC"`
	TanDescription         null.Float64 `name:"tan_description" bson:"tan_description,omitempty" json:"tanDescription"`
	UnionizedTan           null.Float64 `name:"unionized_tan" bson:"unionized_tan,omitempty" json:"unionizedTan"` // should be confirmed the parameter id
	I                      null.Float64 `name:"i" bson:"i,omitempty" json:"i"`                                    // need to be confirmed the parameter id
	PkaS                   null.Float64 `name:"pka_s" bson:"pka_s,omitempty" json:"pkaS"`                         // need to be confirmed the parameter id
	FractionNh3            null.Float64 `name:"fraction_nh3" bson:"fraction_nh3,omitempty" json:"fractionNh3"`
	FractionNh4            null.Float64 `name:"fraction_nh4" bson:"fraction_nh4,omitempty" json:"fractionNh4"`
	XVolume                null.Float64 `name:"x_volume" bson:"x_volume,omitempty" json:"xVolume"`
	YVolume                null.Float64 `name:"y_volume" bson:"y_volume,omitempty" json:"yVolume"`
	NKmno4                 null.Float64 `name:"n_kmno4" bson:"n_kmno4,omitempty" json:"nKmno4"`
	BEKmn04                null.Float64 `name:"be_kmn04" bson:"be_kmn04,omitempty" json:"beKmn04"`
	SampleVolumeTitration  null.Float64 `name:"sample_volume_titration" bson:"sample_volume_titration,omitempty" json:"sampleVolumeTitration"`
	CuVolume               null.Float64 `name:"cu_volume" bson:"cu_volume,omitempty" json:"cuVolume"`
	VEdtaA                 null.Float64 `name:"v_edta_a" bson:"v_edta_a,omitempty" json:"vEdtaA"`
	VEdtaB                 null.Float64 `name:"v_edta_b" bson:"v_edta_b,omitempty" json:"vEdtaB"`
	MEdta                  null.Float64 `name:"m_edta" bson:"m_edta,omitempty" json:"mEdta"`
	MrCaCO3                null.Float64 `name:"mr_caco3" bson:"mr_caco3,omitempty" json:"mrCaCO3"` // need to be confirmed the parameter id
	ArCa                   null.Float64 `name:"ar_ca" bson:"ar_ca,omitempty" json:"arCa"`          // need to be confirmed the parameter id
	ArMg                   null.Float64 `name:"ar_mg" bson:"ar_mg,omitempty" json:"arMg"`          // need to be confirmed the parameter id

	// Lab Plankton Sheet in Web Table Univer
	Actinastrum         null.Float64 `name:"actinastrum" bson:"actinastrum,omitempty" json:"actinastrum"`
	Ankistrodesmus      null.Float64 `name:"ankistrodesmus" bson:"ankistrodesmus,omitempty" json:"ankistrodesmus"`
	Asterococcus        null.Float64 `name:"asterococcus" bson:"asterococcus,omitempty" json:"asterococcus"`
	Carteria            null.Float64 `name:"carteria" bson:"carteria,omitempty" json:"carteria"`
	Chlamydomonos       null.Float64 `name:"chlamydomonos" bson:"chlamydomonos,omitempty" json:"chlamydomonos"`
	Chlorella           null.Float64 `name:"chlorella" bson:"chlorella,omitempty" json:"chlorella"`
	Chloroccoccum       null.Float64 `name:"chloroccoccum" bson:"chloroccoccum,omitempty" json:"chloroccoccum"`
	Chodatella          null.Float64 `name:"chodatella" bson:"chodatella,omitempty" json:"chodatella"`
	Coelastrum          null.Float64 `name:"coelastrum" bson:"coelastrum,omitempty" json:"coelastrum"`
	Cosmarium           null.Float64 `name:"cosmarium" bson:"cosmarium,omitempty" json:"cosmarium"`
	Chlorogonium        null.Float64 `name:"chlorogonium" bson:"chlorogonium,omitempty" json:"chlorogonium"`
	Closterium          null.Float64 `name:"closterium" bson:"closterium,omitempty" json:"closterium"`
	Crucigenia          null.Float64 `name:"crucigenia" bson:"crucigenia,omitempty" json:"crucigenia"`
	Dictyosphaerium     null.Float64 `name:"dictyosphaerium" bson:"dictyosphaerium,omitempty" json:"dictyosphaerium"`
	Desmidium           null.Float64 `name:"desmidium" bson:"desmidium,omitempty" json:"desmidium"`
	Dunaliella          null.Float64 `name:"dunaliella" bson:"dunaliella,omitempty" json:"dunaliella"`
	Echinosphaerella    null.Float64 `name:"echinosphaerella" bson:"echinosphaerella,omitempty" json:"echinosphaerella"`
	Gloeocystis         null.Float64 `name:"gloeocystis" bson:"gloeocystis,omitempty" json:"gloeocystis"`
	Kirchneriella       null.Float64 `name:"kirchneriella" bson:"kirchneriella,omitempty" json:"kirchneriella"`
	Micractinium        null.Float64 `name:"micractinium" bson:"micractinium,omitempty" json:"micractinium"`
	Micrasterias        null.Float64 `name:"micrasterias" bson:"micrasterias,omitempty" json:"micrasterias"`
	Monorapidium        null.Float64 `name:"monorapidium" bson:"monorapidium,omitempty" json:"monorapidium"`
	Nannochloropsis     null.Float64 `name:"nannochloropsis" bson:"nannochloropsis,omitempty" json:"nannochloropsis"`
	Oocystis            null.Float64 `name:"oocystis" bson:"oocystis,omitempty" json:"oocystis"`
	Pandorina           null.Float64 `name:"pandorina" bson:"pandorina,omitempty" json:"pandorina"`
	Planktosphaeria     null.Float64 `name:"planktosphaeria" bson:"planktosphaeria,omitempty" json:"planktosphaeria"`
	Pediastrum          null.Float64 `name:"pediastrum" bson:"pediastrum,omitempty" json:"pediastrum"`
	Pyramimonas         null.Float64 `name:"pyramimonas" bson:"pyramimonas,omitempty" json:"pyramimonas"`
	Selenastrum         null.Float64 `name:"selenastrum" bson:"selenastrum,omitempty" json:"selenastrum"`
	Scenedesmus         null.Float64 `name:"scenedesmus" bson:"scenedesmus,omitempty" json:"scenedesmus"`
	Spinogyra           null.Float64 `name:"spinogyra" bson:"spinogyra,omitempty" json:"spinogyra"`
	Staurastrum         null.Float64 `name:"staurastrum" bson:"staurastrum,omitempty" json:"staurastrum"`
	Tetraedron          null.Float64 `name:"tetraedron" bson:"tetraedron,omitempty" json:"tetraedron"`
	Tetraselmis         null.Float64 `name:"tetraselmis" bson:"tetraselmis,omitempty" json:"tetraselmis"`
	Tribonema           null.Float64 `name:"tribonema" bson:"tribonema,omitempty" json:"tribonema"`
	Treubaria           null.Float64 `name:"treubaria" bson:"treubaria,omitempty" json:"treubaria"`
	Volvox              null.Float64 `name:"volvox" bson:"volvox,omitempty" json:"volvox"`
	Westella            null.Float64 `name:"westella" bson:"westella,omitempty" json:"westella"`
	OtherGreenAlgae     null.Float64 `name:"other_green_algae" bson:"other_green_algae,omitempty" json:"otherGreenAlgae"`
	Anabaena            null.Float64 `name:"anabaena" bson:"anabaena,omitempty" json:"anabaena"`
	Anabaenopsis        null.Float64 `name:"anabaenopsis" bson:"anabaenopsis,omitempty" json:"anabaenopsis"`
	Anacystis           null.Float64 `name:"anacystis" bson:"anacystis,omitempty" json:"anacystis"`
	Aphanochapsa        null.Float64 `name:"aphanochapsa" bson:"aphanochapsa,omitempty" json:"aphanochapsa"`
	Aphanothece         null.Float64 `name:"aphanothece" bson:"aphanothece,omitempty" json:"aphanothece"`
	Chroococus          null.Float64 `name:"chroococus" bson:"chroococus,omitempty" json:"chroococus"`
	Coleosphaerium      null.Float64 `name:"coleosphaerium" bson:"coleosphaerium,omitempty" json:"coleosphaerium"`
	Gloeocapsa          null.Float64 `name:"gloeocapsa" bson:"gloeocapsa,omitempty" json:"gloeocapsa"`
	Gomphospaeria       null.Float64 `name:"gomphospaeria" bson:"gomphospaeria,omitempty" json:"gomphospaeria"`
	Lyngbya             null.Float64 `name:"lyngbya" bson:"lyngbya,omitempty" json:"lyngbya"`
	Merismopedia        null.Float64 `name:"merismopedia" bson:"merismopedia,omitempty" json:"merismopedia"`
	Microcystis         null.Float64 `name:"microcystis" bson:"microcystis,omitempty" json:"microcystis"`
	Oscillatoria        null.Float64 `name:"oscillatoria" bson:"oscillatoria,omitempty" json:"oscillatoria"`
	Pesudoanabayna      null.Float64 `name:"pesudoanabayna" bson:"pesudoanabayna,omitempty" json:"pesudoanabayna"`
	Spirulina           null.Float64 `name:"spirulina" bson:"spirulina,omitempty" json:"spirulina"`
	OtherBlueGreenAlgae null.Float64 `name:"other_blue_green_algae" bson:"other_blue_green_algae,omitempty" json:"otherBlueGreenAlgae"`
	Achnaetes           null.Float64 `name:"achnaetes" bson:"achnaetes,omitempty" json:"achnaetes"`
	Amphipora           null.Float64 `name:"amphipora" bson:"amphipora,omitempty" json:"amphipora"`
	Amphora             null.Float64 `name:"amphora" bson:"amphora,omitempty" json:"amphora"`
	Biddulphia          null.Float64 `name:"biddulphia" bson:"biddulphia,omitempty" json:"biddulphia"`
	Ceratulina          null.Float64 `name:"ceratulina" bson:"ceratulina,omitempty" json:"ceratulina"`
	Chaetocerus         null.Float64 `name:"chaetocerus" bson:"chaetocerus,omitempty" json:"chaetocerus"`
	Cocconeis           null.Float64 `name:"cocconeis" bson:"cocconeis,omitempty" json:"cocconeis"`
	Coscinodiscus       null.Float64 `name:"coscinodiscus" bson:"coscinodiscus,omitempty" json:"coscinodiscus"`
	Cyclotella          null.Float64 `name:"cyclotella" bson:"cyclotella,omitempty" json:"cyclotella"`
	Cylindropyxis       null.Float64 `name:"cylindropyxis" bson:"cylindropyxis,omitempty" json:"cylindropyxis"`
	Cymbella            null.Float64 `name:"cymbella" bson:"cymbella,omitempty" json:"cymbella"`
	Diploneis           null.Float64 `name:"diploneis" bson:"diploneis,omitempty" json:"diploneis"`
	Ditylum             null.Float64 `name:"ditylum" bson:"ditylum,omitempty" json:"ditylum"`
	Gomphonema          null.Float64 `name:"gomphonema" bson:"gomphonema,omitempty" json:"gomphonema"`
	Guinardia           null.Float64 `name:"guinardia" bson:"guinardia,omitempty" json:"guinardia"`
	Gyrosigma           null.Float64 `name:"gyrosigma" bson:"gyrosigma,omitempty" json:"gyrosigma"`
	Melosria            null.Float64 `name:"melosria" bson:"melosria,omitempty" json:"melosria"`
	Navicula            null.Float64 `name:"navicula" bson:"navicula,omitempty" json:"navicula"`
	Nitzchia            null.Float64 `name:"nitzchia" bson:"nitzchia,omitempty" json:"nitzchia"`
	Odontella           null.Float64 `name:"odontella" bson:"odontella,omitempty" json:"odontella"`
	Pinularia           null.Float64 `name:"pinularia" bson:"pinularia,omitempty" json:"pinularia"`
	Pleurosigma         null.Float64 `name:"pleurosigma" bson:"pleurosigma,omitempty" json:"pleurosigma"`
	Rhizosolenia        null.Float64 `name:"rhizosolenia" bson:"rhizosolenia,omitempty" json:"rhizosolenia"`
	Skeletonema         null.Float64 `name:"skeletonema" bson:"skeletonema,omitempty" json:"skeletonema"`
	Streptotecha        null.Float64 `name:"streptotecha" bson:"streptotecha,omitempty" json:"streptotecha"`
	Tabellaria          null.Float64 `name:"tabellaria" bson:"tabellaria,omitempty" json:"tabellaria"`
	Thalassiosira       null.Float64 `name:"thalassiosira" bson:"thalassiosira,omitempty" json:"thalassiosira"`
	OtherChrisophyta    null.Float64 `name:"other_chrisophyta" bson:"other_chrisophyta,omitempty" json:"otherChrisophyta"`
	Astasia             null.Float64 `name:"astasia" bson:"astasia,omitempty" json:"astasia"`
	Cryptoglena         null.Float64 `name:"cryptoglena" bson:"cryptoglena,omitempty" json:"cryptoglena"`
	Euglena             null.Float64 `name:"euglena" bson:"euglena,omitempty" json:"euglena"`
	Trachelomonas       null.Float64 `name:"trachelomonas" bson:"trachelomonas,omitempty" json:"trachelomonas"`
	Phacus              null.Float64 `name:"phacus" bson:"phacus,omitempty" json:"phacus"`
	Rynchomonas         null.Float64 `name:"rynchomonas" bson:"rynchomonas,omitempty" json:"rynchomonas"`
	OtherEuglenophyta   null.Float64 `name:"other_euglenophyta" bson:"other_euglenophyta,omitempty" json:"otherEuglenophyta"`
	Alexandrium         null.Float64 `name:"alexandrium" bson:"alexandrium,omitempty" json:"alexandrium"`
	Gonyaulax           null.Float64 `name:"gonyaulax" bson:"gonyaulax,omitempty" json:"gonyaulax"`
	Gonyostomum         null.Float64 `name:"gonyostomum" bson:"gonyostomum,omitempty" json:"gonyostomum"`
	Gymnodinium         null.Float64 `name:"gymnodinium" bson:"gymnodinium,omitempty" json:"gymnodinium"`
	Gyrodinium          null.Float64 `name:"gyrodinium" bson:"gyrodinium,omitempty" json:"gyrodinium"`
	Noctilluca          null.Float64 `name:"noctilluca" bson:"noctilluca,omitempty" json:"noctilluca"`
	Peridinium          null.Float64 `name:"peridinium" bson:"peridinium,omitempty" json:"peridinium"`
	Protoceraterium     null.Float64 `name:"protoceraterium" bson:"protoceraterium,omitempty" json:"protoceraterium"`
	Prorocentrum        null.Float64 `name:"prorocentrum" bson:"prorocentrum,omitempty" json:"prorocentrum"`
	Protoperidinium     null.Float64 `name:"protoperidinium" bson:"protoperidinium,omitempty" json:"protoperidinium"`
	Symbiodinium        null.Float64 `name:"symbiodinium" bson:"symbiodinium,omitempty" json:"symbiodinium"`
	OtherDinoflagellata null.Float64 `name:"other_dinoflagellata" bson:"other_dinoflagellata,omitempty" json:"otherDinoflagellata"`
	Acanthocystis       null.Float64 `name:"acanthocystis" bson:"acanthocystis,omitempty" json:"acanthocystis"`
	Actinophrys         null.Float64 `name:"actinophrys" bson:"actinophrys,omitempty" json:"actinophrys"`
	Alona               null.Float64 `name:"alona" bson:"alona,omitempty" json:"alona"`
	Amoeba              null.Float64 `name:"amoeba" bson:"amoeba,omitempty" json:"amoeba"`
	Amphileptus         null.Float64 `name:"amphileptus" bson:"amphileptus,omitempty" json:"amphileptus"`
	Brachionus          null.Float64 `name:"brachionus" bson:"brachionus,omitempty" json:"brachionus"`
	Coleps              null.Float64 `name:"coleps" bson:"coleps,omitempty" json:"coleps"`
	Cyclidium           null.Float64 `name:"cyclidium" bson:"cyclidium,omitempty" json:"cyclidium"`
	Entosiphon          null.Float64 `name:"entosiphon" bson:"entosiphon,omitempty" json:"entosiphon"`
	Epiphanes           null.Float64 `name:"epiphanes" bson:"epiphanes,omitempty" json:"epiphanes"`
	Euplotes            null.Float64 `name:"euplotes" bson:"euplotes,omitempty" json:"euplotes"`
	Mesodinium          null.Float64 `name:"mesodinium" bson:"mesodinium,omitempty" json:"mesodinium"`
	Paramecium          null.Float64 `name:"paramecium" bson:"paramecium,omitempty" json:"paramecium"`
	Stentor             null.Float64 `name:"stentor" bson:"stentor,omitempty" json:"stentor"`
	Strombidinopsis     null.Float64 `name:"strombidinopsis" bson:"strombidinopsis,omitempty" json:"strombidinopsis"`
	Strombidinium       null.Float64 `name:"strombidinium" bson:"strombidinium,omitempty" json:"strombidinium"`
	Tintinopsis         null.Float64 `name:"tintinopsis" bson:"tintinopsis,omitempty" json:"tintinopsis"`
	Trichocerca         null.Float64 `name:"trichocerca" bson:"trichocerca,omitempty" json:"trichocerca"`
	Vorticella          null.Float64 `name:"vorticella" bson:"vorticella,omitempty" json:"vorticella"`
	Zoothamnium         null.Float64 `name:"zoothamnium" bson:"zoothamnium,omitempty" json:"zoothamnium"`
	OtherZooplankton    null.Float64 `name:"other_zooplankton" bson:"other_zooplankton,omitempty" json:"otherZooplankton"`
	Cryptomonas         null.Float64 `name:"cryptomonas" bson:"cryptomonas,omitempty" json:"cryptomonas"`
	Prymenesium         null.Float64 `name:"prymenesium" bson:"prymenesium,omitempty" json:"prymenesium"`
	Isochrysis          null.Float64 `name:"isochrysis" bson:"isochrysis,omitempty" json:"isochrysis"`
	OtherPlankton       null.Float64 `name:"other_plankton" bson:"other_plankton,omitempty" json:"other_plankton"`
	BacteryTotal        null.Float64 `name:"bactery_total" bson:"bactery_total,omitempty" json:"bacteryTotal"`

	// Bactery TVC and TBC (In Lab Calculation Sheet)
	InoculumTvc1            null.Float64 `name:"inoculum_tvc_1" bson:"inoculumTvc1,omitempty" json:"inoculumTvc1"`
	DilutionFactorTvc1      null.Float64 `name:"dilution_factor_tvc_1" bson:"dilutionFactorTvc1,omitempty" json:"dilutionFactorTvc1"`
	YellowPetri1Tvc1        null.Float64 `name:"yellow_petri1_tvc_1" bson:"yellowPetri1Tvc1,omitempty" json:"yellowPetri1Tvc1"`
	GreenPetri1Tvc1         null.Float64 `name:"green_petri1_tvc_1" bson:"greenPetri1Tvc1,omitempty" json:"greenPetri1Tvc1"`
	BlackPetri1Tvc1         null.Float64 `name:"black_petri1_tvc_1" bson:"blackPetri1Tvc1,omitempty" json:"blackPetri1Tvc1"`
	AdjustedCountPetri1Tvc1 null.Float64 `name:"adjusted_count_petri1_tvc_1" bson:"adjustedCountPetri1Tvc1,omitempty" json:"adjustedCountPetri1Tvc1"`
	YellowPetri2Tvc1        null.Float64 `name:"yellow_petri2_tvc_1" bson:"yellowPetri2Tvc1,omitempty" json:"yellowPetri2Tvc1"`
	GreenPetri2Tvc1         null.Float64 `name:"green_petri2_tvc_1" bson:"greenPetri2Tvc1,omitempty" json:"greenPetri2Tvc1"`
	BlackPetri2Tvc1         null.Float64 `name:"black_petri2_tvc_1" bson:"blackPetri2Tvc1,omitempty" json:"blackPetri2Tvc1"`
	AdjustedCountPetri2Tvc1 null.Float64 `name:"adjusted_count_petri2_tvc_1" bson:"adjustedCountPetri2Tvc1,omitempty" json:"adjustedCountPetri2Tvc1"`
	TotalCountTvc1          null.Float64 `name:"total_count_tvc_1" bson:"totalCountTvc1,omitempty" json:"totalCountTvc1"`
	DilutionFactorTvc2      null.Float64 `name:"dilution_factor_tvc_2" bson:"dilutionFactorTvc2,omitempty" json:"dilutionFactorTvc2"`
	YellowPetri1Tvc2        null.Float64 `name:"yellow_petri1_tvc_2" bson:"yellowPetri1Tvc2,omitempty" json:"yellowPetri1Tvc2"`
	GreenPetri1Tvc2         null.Float64 `name:"green_petri1_tvc_2" bson:"greenPetri1Tvc2,omitempty" json:"greenPetri1Tvc2"`
	BlackPetri1Tvc2         null.Float64 `name:"black_petri1_tvc_2" bson:"blackPetri1Tvc2,omitempty" json:"blackPetri1Tvc2"`
	AdjustedCountPetri1Tvc2 null.Float64 `name:"adjusted_count_petri1_tvc_2" bson:"adjustedCountPetri1Tvc2,omitempty" json:"adjustedCountPetri1Tvc2"`
	TotalCountTvc2          null.Float64 `name:"total_count_tvc_2" bson:"totalCountTvc2,omitempty" json:"totalCountTvc2"`
	InoculumTbc1            null.Float64 `name:"inoculum_tbc_1" bson:"inoculumTbc1,omitempty" json:"inoculumTbc1"`
	DilutionFactorTbc1      null.Float64 `name:"dilution_factor_tbc_1" bson:"dilutionFactorTbc1,omitempty" json:"dilutionFactorTbc1"`
	Petri1Tbc1              null.Float64 `name:"petri1_tbc_1" bson:"petri1Tbc1,omitempty" json:"petri1Tbc1"`
	AdjustedCountPetri1Tbc1 null.Float64 `name:"adjusted_count_petri1_tbc_1" bson:"adjustedCountPetri1Tbc1,omitempty" json:"adjustedCountPetri1Tbc1"`
	Petri2Tbc1              null.Float64 `name:"petri2_tbc_1" bson:"petri2Tbc1,omitempty" json:"petri2Tbc1"`
	AdjustedCountPetri2Tbc1 null.Float64 `name:"adjusted_count_petri2_tbc_1" bson:"adjustedCountPetri2Tbc1,omitempty" json:"adjustedCountPetri2Tbc1"`
	TotalCountTbc1          null.Float64 `name:"total_count_tbc_1" bson:"totalCountTbc1,omitempty" json:"totalCountTbc1"`
	DilutionFactorTbc2      null.Float64 `name:"dilution_factor_tbc_2" bson:"dilutionFactorTbc2,omitempty" json:"dilutionFactorTbc2"`
	Petri1Tbc2              null.Float64 `name:"petri1_tbc_2" bson:"petri1Tbc2,omitempty" json:"petri1Tbc2"`
	AdjustedCountPetri1Tbc2 null.Float64 `name:"adjusted_count_petri1_tbc_2" bson:"adjustedCountPetri1Tbc2,omitempty" json:"adjustedCountPetri1Tbc2"`
	Petri2Tbc2              null.Float64 `name:"petri2_tbc_2" bson:"petri2Tbc2,omitempty" json:"petri2Tbc2"`
	AdjustedCountPetri2Tbc2 null.Float64 `name:"adjusted_count_petri2_tbc_2" bson:"adjustedCountPetri2Tbc2,omitempty" json:"adjustedCountPetri2Tbc2"`
	TotalCountTbc2          null.Float64 `name:"total_count_tbc_2" bson:"totalCountTbc2,omitempty" json:"totalCountTbc2"`

	// Calibration Curvature
	NitritSlope      null.Float64 `name:"nitrit_slope" bson:"nitrit_slope,omitempty" json:"nitritSlope"`
	NitritIntercept  null.Float64 `name:"nitrit_intercept" bson:"nitrit_intercept,omitempty" json:"nitritIntercept"`
	PhospatSlope     null.Float64 `name:"phospat_slope" bson:"phospat_slope,omitempty" json:"phospatSlope"`
	PhospatIntercept null.Float64 `name:"phospat_intercept" bson:"phospat_intercept,omitempty" json:"phospatIntercept"`
	TANSlope         null.Float64 `name:"tan_slope" bson:"tan_slope,omitempty" json:"tanSlope"`
	TANIntercept     null.Float64 `name:"tan_intercept" bson:"tan_intercept,omitempty" json:"tanIntercept"`

	// Lab Analyst Parameter End

	// Additional Attributes
	Status    int64     `bson:"status,omitempty" json:"status" hideMetrics:"true"`
	CreatedAt null.Time `bson:"created_at,omitempty" json:"createdAt" hideMetrics:"true" swaggertype:"string" example:"2022-06-21T10:32:29Z"`
	CreatedBy string    `bson:"created_by,omitempty" json:"createdBy" hideMetrics:"true" swaggertype:"string"`
	UpdatedAt null.Time `bson:"updated_at,omitempty" json:"updatedAt" hideMetrics:"true" swaggertype:"string" example:"2022-06-21T10:32:29Z"`
	UpdatedBy string    `bson:"updated_by,omitempty" json:"updatedBy" hideMetrics:"true" swaggertype:"string"`
	DeletedAt null.Time `bson:"deleted_at,omitempty" json:"deletedAt" hideMetrics:"true" swaggertype:"string" example:"2022-06-21T10:32:29Z"`
	DeletedBy string    `bson:"deleted_by,omitempty" json:"deletedBy" hideMetrics:"true" swaggertype:"string"`
}

type AgGrids []AgGridMetric
