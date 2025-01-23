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
	FarmID       null.Int64  `name:"farm_id" json:"farm_id" bson:"farm_id,omitempty"`
	CycleID      null.Int64  `name:"cycle_id" json:"cycle_id" bson:"cycle_id,omitempty"`
	PondID       null.Int64  `name:"pond_id" json:"pond_id" bson:"pond_id,omitempty"`
	TemplateType null.String `name:"template_type" json:"template_type" bson:"template_type,omitempty"`
	DocOrDop     null.String `name:"doc_or_dop" json:"doc_or_dop" bson:"doc_or_dop,omitempty"`

	// Y Axis
	YAxis null.Int64 `name:"y_axis" json:"y_axis" bson:"y_axis,omitempty"`

	// General Attributes
	MetricDate     null.String  `name:"metric_date" json:"metric_date" bson:"metric_date,omitempty"`
	PondName       null.String  `name:"pond_name" json:"pond_name" bson:"pond_name,omitempty"`
	StockingActual null.Float64 `name:"stocking_actual" json:"stocking_actual" bson:"stocking_actual,omitempty"`
	Doc            null.Int64   `name:"doc" json:"doc" bson:"doc,omitempty"`
	Dop            null.Int64   `name:"dop" json:"dop" bson:"dop,omitempty"`
	DayNumber      null.Float64 `name:"day_number" json:"day_number" bson:"day_number,omitempty"`

	// Farm Technician Parameters
	WaterHeightMorning             null.Float64 `name:"water_height_morning" bson:"water_height_morning,omitempty" json:"water_height_morning"`
	WaterHeight                    null.Float64 `name:"water_height" bson:"water_height,omitempty" json:"water_height"`
	WaterColorMorning              null.String  `name:"water_color_morning" bson:"water_color_morning,omitempty" json:"water_color_morning"`
	WaterColorEvening              null.String  `name:"water_color_evening" bson:"water_color_evening,omitempty" json:"water_color_evening"`
	BrightnessMorning              null.Float64 `name:"brightness_morning" bson:"brightness_morning,omitempty" json:"brightness_morning"`
	BrightnessEvening              null.Float64 `name:"brightness_evening" bson:"brightness_evening,omitempty" json:"brightness_evening"`
	WeatherMorning                 null.String  `name:"weather_morning" bson:"weather_morning,omitempty" json:"weather_morning"`
	WeatherEvening                 null.String  `name:"weather_evening" bson:"weather_evening,omitempty" json:"weather_evening"`
	WeatherNight                   null.String  `name:"weather_night" bson:"weather_night,omitempty" json:"weather_night"`
	SyphonHourMorning              null.Float64 `name:"syphon_hour_morning" bson:"syphon_hour_morning,omitempty" json:"syphon_hour_morning"`
	SyphonHourEvening              null.Float64 `name:"syphon_hour_evening" bson:"syphon_hour_evening,omitempty" json:"syphon_hour_evening"`
	WaterOut                       null.Float64 `name:"water_out" bson:"water_out,omitempty" json:"water_out"`
	WaterIn                        null.Float64 `name:"water_in" bson:"water_in,omitempty" json:"water_in"`
	WaterOverflowDuration          null.Float64 `name:"water_overflow_duration" bson:"water_overflow_duration,omitempty" json:"water_overflow_duration"`
	MortalityDaily                 null.Float64 `name:"mortality_daily" bson:"mortality_daily,omitempty" json:"mortality_daily"`
	MortalityCum                   null.Float64 `name:"mortality_cum" bson:"mortality_cum,omitempty" json:"mortality_cum"`
	PartialHarvest                 null.Float64 `name:"partial_harvest" bson:"partial_harvest,omitempty" json:"partial_harvest"`
	SrTodayPercentage              null.Float64 `name:"sr_today_percentage" bson:"sr_today_percentage,omitempty" json:"sr_today_percentage"`
	WaterWheel1HpMorning           null.Float64 `name:"water_wheel_1hp_morning" bson:"water_wheel_1hp_morning,omitempty" json:"water_wheel_1hp_morning"`
	WaterWheel1HpMorningDuration   null.Float64 `name:"water_wheel_1hp_morning_duration" bson:"water_wheel_1hp_morning_duration,omitempty" json:"water_wheel_1hp_morning_duration"`
	WaterWheel1HpAfternoon         null.Float64 `name:"water_wheel_1hp_afternoon" bson:"water_wheel_1hp_afternoon,omitempty" json:"water_wheel_1hp_afternoon"`
	WaterWheel1HpAfternoonDuration null.Float64 `name:"water_wheel_1hp_afternoon_duration" bson:"water_wheel_1hp_afternoon_duration,omitempty" json:"water_wheel_1hp_afternoon_duration"`
	WaterWheel2HpMorning           null.Float64 `name:"water_wheel_2hp_morning" bson:"water_wheel_2hp_morning,omitempty" json:"water_wheel_2hp_morning"`
	WaterWheel2HpMorningDuration   null.Float64 `name:"water_wheel_2hp_morning_duration" bson:"water_wheel_2hp_morning_duration,omitempty" json:"water_wheel_2hp_morning_duration"`
	WaterWheel2HpAfternoon         null.Float64 `name:"water_wheel_2hp_afternoon" bson:"water_wheel_2hp_afternoon,omitempty" json:"water_wheel_2hp_afternoon"`
	WaterWheel2HpAfternoonDuration null.Float64 `name:"water_wheel_2hp_afternoon_duration" bson:"water_wheel_2hp_afternoon_duration,omitempty" json:"water_wheel_2hp_afternoon_duration"`
	BlowerMorning                  null.Float64 `name:"blower_morning" bson:"blower_morning,omitempty" json:"blower_morning"`
	BlowerMorningDuration          null.Float64 `name:"blower_morning_duration" bson:"blower_morning_duration,omitempty" json:"blower_morning_duration"`
	BlowerEvening                  null.Float64 `name:"blower_evening" bson:"blower_evening,omitempty" json:"blower_evening"`
	BlowerAfternoonDuration        null.Float64 `name:"blower_afternoon_duration" bson:"blower_afternoon_duration,omitempty" json:"blower_afternoon_duration"`
	TotalPowerMorning              null.Float64 `name:"total_power_morning" bson:"total_power_morning,omitempty" json:"total_power_morning"`
	TotalPowerEvening              null.Float64 `name:"total_power_evening" bson:"total_power_evening,omitempty" json:"total_power_evening"`
	BrandFeed                      null.String  `name:"brand_feed" bson:"brand_feed,omitempty" json:"brand_feed"`
	FeedCode                       null.String  `name:"feed_code" bson:"feed_code,omitempty" json:"feed_code"`
	FeedSession1Weight             null.Float64 `name:"feed_session1_weight" bson:"feed_session1_weight,omitempty" json:"feed_session1_weight"`
	FeedSession2Weight             null.Float64 `name:"feed_session2_weight" bson:"feed_session2_weight,omitempty" json:"feed_session2_weight"`
	FeedSession3Weight             null.Float64 `name:"feed_session3_weight" bson:"feed_session3_weight,omitempty" json:"feed_session3_weight"`
	FeedSession4Weight             null.Float64 `name:"feed_session4_weight" bson:"feed_session4_weight,omitempty" json:"feed_session4_weight"`
	FeedSession5Weight             null.Float64 `name:"feed_session5_weight" bson:"feed_session5_weight,omitempty" json:"feed_session5_weight"`
	FeedSession6Weight             null.Float64 `name:"feed_session6_weight" bson:"feed_session6_weight,omitempty" json:"feed_session6_weight"`
	BrandFeedMix                   null.String  `name:"brand_feed_mix" bson:"brand_feed_mix,omitempty" json:"brand_feed_mix"`
	FeedCodeMix                    null.String  `name:"brand_feed_code_mix" bson:"brand_feed_code_mix,omitempty" json:"brand_feed_code_mix"`
	FeedSession1WeightMix          null.Float64 `name:"feed_session1_weight_mix" bson:"feed_session1_weight_mix,omitempty" json:"feed_session1_weight_mix"`
	FeedSession2WeightMix          null.Float64 `name:"feed_session2_weight_mix" bson:"feed_session2_weight_mix,omitempty" json:"feed_session2_weight_mix"`
	FeedSession3WeightMix          null.Float64 `name:"feed_session3_weight_mix" bson:"feed_session3_weight_mix,omitempty" json:"feed_session3_weight_mix"`
	FeedSession4WeightMix          null.Float64 `name:"feed_session4_weight_mix" bson:"feed_session4_weight_mix,omitempty" json:"feed_session4_weight_mix"`
	FeedSession5WeightMix          null.Float64 `name:"feed_session5_weight_mix" bson:"feed_session5_weight_mix,omitempty" json:"feed_session5_weight_mix"`
	FeedSession6WeightMix          null.Float64 `name:"feed_session6_weight_mix" bson:"feed_session6_weight_mix,omitempty" json:"feed_session6_weight_mix"`
	TotalFeedActual                null.Float64 `name:"total_feed_actual" bson:"total_feed_actual,omitempty" json:"total_feed_actual"`
	TotalCumFeed                   null.Float64 `name:"total_cum_feed" bson:"total_cum_feed,omitempty" json:"total_cum_feed"`
	FeedNote                       null.String  `name:"feed_note_session1" bson:"feed_note_session1,omitempty" json:"feed_note_session1"`
	ShrimpCondition                null.String  `name:"shrimp_condition_session1" bson:"shrimp_condition_session1,omitempty" json:"shrimp_condition_session1"`
	AncoPercentage                 null.Float64 `name:"anco_percentage" bson:"anco_percentage,omitempty" json:"anco_percentage"`
	AncoSess1Anco1                 null.String  `name:"anco_session1_anco1" bson:"anco_session1_anco1,omitempty" json:"anco_session1_anco1"`
	AncoSess1Anco2                 null.String  `name:"anco_session1_anco2" bson:"anco_session1_anco2,omitempty" json:"anco_session1_anco2"`
	AncoSess1Anco3                 null.String  `name:"anco_session1_anco3" bson:"anco_session1_anco3,omitempty" json:"anco_session1_anco3"`
	AncoSess1Anco4                 null.String  `name:"anco_session1_anco4" bson:"anco_session1_anco4,omitempty" json:"anco_session1_anco4"`
	AncoSess1Anco5                 null.String  `name:"anco_session1_anco5" bson:"anco_session1_anco5,omitempty" json:"anco_session1_anco5"`
	AncoSess2Anco1                 null.String  `name:"anco_session2_anco1" bson:"anco_session2_anco1,omitempty" json:"anco_session2_anco1"`
	AncoSess2Anco2                 null.String  `name:"anco_session2_anco2" bson:"anco_session2_anco2,omitempty" json:"anco_session2_anco2"`
	AncoSess2Anco3                 null.String  `name:"anco_session2_anco3" bson:"anco_session2_anco3,omitempty" json:"anco_session2_anco3"`
	AncoSess2Anco4                 null.String  `name:"anco_session2_anco4" bson:"anco_session2_anco4,omitempty" json:"anco_session2_anco4"`
	AncoSess2Anco5                 null.String  `name:"anco_session2_anco5" bson:"anco_session2_anco5,omitempty" json:"anco_session2_anco5"`
	AncoSess3Anco1                 null.String  `name:"anco_session3_anco1" bson:"anco_session3_anco1,omitempty" json:"anco_session3_anco1"`
	AncoSess3Anco2                 null.String  `name:"anco_session3_anco2" bson:"anco_session3_anco2,omitempty" json:"anco_session3_anco2"`
	AncoSess3Anco3                 null.String  `name:"anco_session3_anco3" bson:"anco_session3_anco3,omitempty" json:"anco_session3_anco3"`
	AncoSess3Anco4                 null.String  `name:"anco_session3_anco4" bson:"anco_session3_anco4,omitempty" json:"anco_session3_anco4"`
	AncoSess3Anco5                 null.String  `name:"anco_session3_anco5" bson:"anco_session3_anco5,omitempty" json:"anco_session3_anco5"`
	Treatment1Brand                null.String  `name:"treatment1_brand" bson:"treatment_1_brand,omitempty" json:"treatment1_brand"`
	Treatment1Type                 null.String  `name:"treatment1_type" bson:"treatment_1_type,omitempty" json:"treatment1_type"`
	Treatment1Total                null.Float64 `name:"treatment1_total" bson:"treatment_1_total,omitempty" json:"treatment1_total"`
	Treatment1Unit                 null.String  `name:"treatment1_unit" bson:"treatment_1_unit,omitempty" json:"treatment1_unit"`
	Treatment2Brand                null.String  `name:"treatment2_brand" bson:"treatment_2_brand,omitempty" json:"treatment2_brand"`
	Treatment2Type                 null.String  `name:"treatment2_type" bson:"treatment_2_type,omitempty" json:"treatment2_type"`
	Treatment2Total                null.Float64 `name:"treatment2_total" bson:"treatment_2_total,omitempty" json:"treatment2_total"`
	Treatment2Unit                 null.String  `name:"treatment2_unit" bson:"treatment_2_unit,omitempty" json:"treatment2_unit"`
	Treatment3Brand                null.String  `name:"treatment3_brand" bson:"treatment_3_brand,omitempty" json:"treatment3_brand"`
	Treatment3Type                 null.String  `name:"treatment3_type" bson:"treatment_3_type,omitempty" json:"treatment3_type"`
	Treatment3Total                null.Float64 `name:"treatment3_total" bson:"treatment_3_total,omitempty" json:"treatment3_total"`
	Treatment3Unit                 null.String  `name:"treatment3_unit" bson:"treatment_3_unit,omitempty" json:"treatment3_unit"`
	Treatment4Brand                null.String  `name:"treatment4_brand" bson:"treatment_4_brand,omitempty" json:"treatment4_brand"`
	Treatment4Type                 null.String  `name:"treatment4_type" bson:"treatment_4_type,omitempty" json:"treatment4_type"`
	Treatment4Total                null.Float64 `name:"treatment4_total" bson:"treatment_4_total,omitempty" json:"treatment4_total"`
	Treatment4Unit                 null.String  `name:"treatment4_unit" bson:"treatment_4_unit,omitempty" json:"treatment4_unit"`
	Treatment5Brand                null.String  `name:"treatment5_brand" bson:"treatment_5_brand,omitempty" json:"treatment5_brand"`
	Treatment5Type                 null.String  `name:"treatment5_type" bson:"treatment_5_type,omitempty" json:"treatment5_type"`
	Treatment5Total                null.Float64 `name:"treatment5_total" bson:"treatment_5_total,omitempty" json:"treatment5_total"`
	Treatment5Unit                 null.String  `name:"treatment5_unit" bson:"treatment_5_unit,omitempty" json:"treatment5_unit"`
	Treatment6Brand                null.String  `name:"treatment6_brand" bson:"treatment_6_brand,omitempty" json:"treatment6_brand"`
	Treatment6Type                 null.String  `name:"treatment6_type" bson:"treatment_6_type,omitempty" json:"treatment6_type"`
	Treatment6Total                null.Float64 `name:"treatment6_total" bson:"treatment_6_total,omitempty" json:"treatment6_total"`
	Treatment6Unit                 null.String  `name:"treatment6_unit" bson:"treatment_6_unit,omitempty" json:"treatment6_unit"`
	Treatment7Brand                null.String  `name:"treatment7_brand" bson:"treatment_7_brand,omitempty" json:"treatment7_brand"`
	Treatment7Type                 null.String  `name:"treatment7_type" bson:"treatment_7_type,omitempty" json:"treatment7_type"`
	Treatment7Total                null.Float64 `name:"treatment7_total" bson:"treatment_7_total,omitempty" json:"treatment7_total"`
	Treatment7Unit                 null.String  `name:"treatment7_unit" bson:"treatment_7_unit,omitempty" json:"treatment7_unit"`
	Treatment8Brand                null.String  `name:"treatment8_brand" bson:"treatment_8_brand,omitempty" json:"treatment8_brand"`
	Treatment8Type                 null.String  `name:"treatment8_type" bson:"treatment_8_type,omitempty" json:"treatment8_type"`
	Treatment8Total                null.Float64 `name:"treatment8_total" bson:"treatment_8_total,omitempty" json:"treatment8_total"`
	Treatment8Unit                 null.String  `name:"treatment8_unit" bson:"treatment_8_unit,omitempty" json:"treatment8_unit"`
	Treatment9Brand                null.String  `name:"treatment9_brand" bson:"treatment_9_brand,omitempty" json:"treatment9_brand"`
	Treatment9Type                 null.String  `name:"treatment9_type" bson:"treatment_9_type,omitempty" json:"treatment9_type"`
	Treatment9Total                null.Float64 `name:"treatment9_total" bson:"treatment_9_total,omitempty" json:"treatment9_total"`
	Treatment9Unit                 null.String  `name:"treatment9_unit" bson:"treatment_9_unit,omitempty" json:"treatment9_unit"`
	Treatment10Brand               null.String  `name:"treatment10_brand" bson:"treatment_10_brand,omitempty" json:"treatment10_brand"`
	Treatment10Type                null.String  `name:"treatment10_type" bson:"treatment_10_type,omitempty" json:"treatment10_type"`
	Treatment10Total               null.Float64 `name:"treatment10_total" bson:"treatment_10_total,omitempty" json:"treatment10_total"`
	Treatment10Unit                null.String  `name:"treatment10_unit" bson:"treatment_10_unit,omitempty" json:"treatment10_unit"`
	TreatmentNote                  null.String  `name:"treatment_note" bson:"treatment_note,omitempty" json:"treatment_note"`

	BrandFeedSession2 null.String `name:"brand_feed_session2" bson:"brand_feed_session_2,omitempty" json:"brand_feed_session2"`
	FeedCodeSession2  null.String `name:"feed_code_session2" bson:"feed_code_session_2,omitempty" json:"feed_code_session2"`
	BrandFeedSession3 null.String `name:"brand_feed_session3" bson:"brand_feed_session_3,omitempty" json:"brand_feed_session3"`
	FeedCodeSession3  null.String `name:"feed_code_session3" bson:"feed_code_session_3,omitempty" json:"feed_code_session3"`
	BrandFeedSession4 null.String `name:"brand_feed_session4" bson:"brand_feed_session_4,omitempty" json:"brand_feed_session4"`
	FeedCodeSession4  null.String `name:"feed_code_session4" bson:"feed_code_session_4,omitempty" json:"feed_code_session4"`
	BrandFeedSession5 null.String `name:"brand_feed_session5" bson:"brand_feed_session_5,omitempty" json:"brand_feed_session5"`
	FeedCodeSession5  null.String `name:"feed_code_session5" bson:"feed_code_session_5,omitempty" json:"feed_code_session5"`
	BrandFeedSession6 null.String `name:"brand_feed_session6" bson:"brand_feed_session_6,omitempty" json:"brand_feed_session6"`
	FeedCodeSession6  null.String `name:"feed_code_session6" bson:"feed_code_session_6,omitempty" json:"feed_code_session6"`

	BrandFeedMixSession2     null.String `name:"brand_feed_mix_session2" bson:"brand_feed_mix_session_2,omitempty" json:"brand_feed_mix_session2"`
	BrandFeedCodeMixSession2 null.String `name:"brand_feed_code_mix_session2" bson:"brand_feed_code_mix_session_2,omitempty" json:"brand_feed_code_mix_session2"`
	BrandFeedMixSession3     null.String `name:"brand_feed_mix_session3" bson:"brand_feed_mix_session_3,omitempty" json:"brand_feed_mix_session3"`
	BrandFeedCodeMixSession3 null.String `name:"brand_feed_code_mix_session3" bson:"brand_feed_code_mix_session_3,omitempty" json:"brand_feed_code_mix_session3"`
	BrandFeedMixSession4     null.String `name:"brand_feed_mix_session4" bson:"brand_feed_mix_session_4,omitempty" json:"brand_feed_mix_session4"`
	BrandFeedCodeMixSession4 null.String `name:"brand_feed_code_mix_session4" bson:"brand_feed_code_mix_session_4,omitempty" json:"brand_feed_code_mix_session4"`
	BrandFeedMixSession5     null.String `name:"brand_feed_mix_session5" bson:"brand_feed_mix_session_5,omitempty" json:"brand_feed_mix_session5"`
	BrandFeedCodeMixSession5 null.String `name:"brand_feed_code_mix_session5" bson:"brand_feed_code_mix_session_5,omitempty" json:"brand_feed_code_mix_session5"`
	BrandFeedMixSession6     null.String `name:"brand_feed_mix_session6" bson:"brand_feed_mix_session_6,omitempty" json:"brand_feed_mix_session6"`
	BrandFeedCodeMixSession6 null.String `name:"brand_feed_code_mix_session6" bson:"brand_feed_code_mix_session_6,omitempty" json:"brand_feed_code_mix_session6"`

	// Farm Technician Parameter End

	// Lab analyst Parameters
	// "Lab Analyst" sheet

	// SAL is already on Farm Technician
	Sal                          null.Float64 `name:"sal" bson:"sal,omitempty" json:"sal"`
	SalTime                      null.String  `name:"salinity_time" bson:"salinity_time,omitempty" json:"salinity_time"`
	DOMorning                    null.Float64 `name:"do_time0500" bson:"do_time0500,omitempty" json:"do_time0500"`
	DOMorningTime                null.String  `name:"do_time0500_time" bson:"do_time0500_time,omitempty" json:"do_time0500_time"`
	DOAfternoon                  null.Float64 `name:"do_time1300" bson:"do_time1300,omitempty" json:"do_time1300"`
	DOAfternoonTime              null.String  `name:"do_time1300_time" bson:"do_time1300_time,omitempty" json:"do_time1300_time"`
	DONight                      null.Float64 `name:"do_time2100" bson:"do_time2100,omitempty" json:"do_time2100"`
	DONightTime                  null.String  `name:"do_time2100_time" bson:"do_time2100_time,omitempty" json:"do_time2100_time"`
	PHMorning                    null.Float64 `name:"ph_time0430" bson:"ph_time0430,omitempty" json:"ph_time0430"`
	PHMorningTime                null.String  `name:"ph_time0430_time" bson:"ph_time0430_time,omitempty" json:"ph_time0430_time"`
	PHAfternoon                  null.Float64 `name:"ph_time1300" bson:"ph_time1300,omitempty" json:"ph_time1300"`
	PHAfternoonTime              null.String  `name:"ph_time1300_time" bson:"ph_time1300_time,omitempty" json:"ph_time1300_time"`
	RangePH                      null.Float64 `name:"range_ph" bson:"range_ph,omitempty" json:"range_ph"`
	TemperatureMorning           null.Float64 `name:"temperature_time0500" bson:"temperature_time0500,omitempty" json:"temperature_time0500"`
	TemperatureMorningTime       null.String  `name:"temperature_time0500_time" bson:"temperature_time0500_time,omitempty" json:"temperature_time0500_time"`
	TemperatureAfternoon         null.Float64 `name:"temperature_time1400" bson:"temperature_time1400,omitempty" json:"temperature_time1400"`
	TemperatureAfternoonTime     null.String  `name:"temperature_time1400_time" bson:"temperature_time1400_time,omitempty" json:"temperature_time1400_time"`
	TemperatureNight             null.Float64 `name:"temperature_time2100" bson:"temperature_time2100,omitempty" json:"temperature_time2100"`
	TemperatureNightTime         null.String  `name:"temperature_time2100_time" bson:"temperature_time2100_time,omitempty" json:"temperature_time2100_time"`
	ORPMorning                   null.Float64 `name:"orp_time0500" bson:"orp_time0500,omitempty" json:"orp_time0500"`
	Tan                          null.Float64 `name:"tan" bson:"tan,omitempty" json:"tan"`
	Nh3                          null.Float64 `name:"nh3" bson:"nh3,omitempty" json:"nh3"`
	Nh4                          null.Float64 `name:"nh4" bson:"nh4,omitempty" json:"nh4"`
	No3                          null.Float64 `name:"no3" bson:"no3,omitempty" json:"no3"`
	No2                          null.Float64 `name:"no2" bson:"no2,omitempty" json:"no2"`
	Po4                          null.Float64 `name:"po4" bson:"po4,omitempty" json:"po4"`
	RatioNP                      null.Float64 `name:"ratio_np" bson:"ratio_np,omitempty" json:"ratio_np"`
	Oh                           null.Float64 `name:"oh" bson:"oh,omitempty" json:"oh"`
	Co3                          null.Float64 `name:"co3" bson:"co3,omitempty" json:"co3"`
	Hco3                         null.Float64 `name:"hco3" bson:"hco3,omitempty" json:"hco3"`
	TAlk                         null.Float64 `name:"t_alk" bson:"t_alk,omitempty" json:"t_alk"`
	Ca                           null.Float64 `name:"ca" bson:"ca,omitempty" json:"ca"`
	Mg                           null.Float64 `name:"mg" bson:"mg,omitempty" json:"mg"`
	MgCa                         null.Float64 `name:"mgca" bson:"mgca,omitempty" json:"mgca"`
	Camg                         null.Float64 `name:"camg" bson:"camg,omitempty" json:"camg"`
	THard                        null.Float64 `name:"t_hard" bson:"t_hard,omitempty" json:"t_hard"`
	Tom                          null.Float64 `name:"tom" bson:"tom,omitempty" json:"tom"`
	BacteryGaPercentage          null.Float64 `name:"bactery_ga_percentage" bson:"bactery_ga_percentage,omitempty" json:"bactery_ga_percentage"`
	BacteryBgaPercentage         null.Float64 `name:"bactery_bga_percentage" bson:"bactery_bga_percentage,omitempty" json:"bactery_bga_percentage"`
	BacteryDiaPercentage         null.Float64 `name:"bactery_dia_percentage" bson:"bactery_dia_percentage,omitempty" json:"bactery_dia_percentage"`
	BacteryEuglenoPercentage     null.Float64 `name:"bactery_eugleno_percentage" bson:"bactery_eugleno_percentage,omitempty" json:"bactery_eugleno_percentage"`
	BacteryDinoPercentage        null.Float64 `name:"bactery_dino_percentage" bson:"bactery_dino_percentage,omitempty" json:"bactery_dino_percentage"`
	BacteryZooPercentage         null.Float64 `name:"bactery_zoo_percentage" bson:"bactery_zoo_percentage,omitempty" json:"bactery_zoo_percentage"`
	TotalOtherPlanktonPercentage null.Float64 `name:"total_other_plankton_percentage" bson:"total_other_plankton_percentage,omitempty" json:"total_other_plankton_percentage"`
	BacteryVibrioY               null.Float64 `name:"bactery_vibrio_y" bson:"bactery_vibrio_y,omitempty" json:"bactery_vibrio_y"`
	BacteryVibrioG               null.Float64 `name:"bactery_vibrio_g" bson:"bactery_vibrio_g,omitempty" json:"bactery_vibrio_g"`
	BacteryBlack                 null.String  `name:"bactery_black" bson:"bactery_black,omitempty" json:"bactery_black"`
	BacteryVibrioTvc             null.Float64 `name:"bactery_vibrio_tvc" bson:"bactery_vibrio_tvc,omitempty" json:"bactery_vibrio_tvc"`
	BacteryTbc                   null.Float64 `name:"bactery_tbc" bson:"bactery_tbc,omitempty" json:"bactery_tbc"`
	TvcTbcPercentage             null.Float64 `name:"tvc_tbc_percentage" bson:"tvc_tbc_percentage,omitempty" json:"tvc_tbc_percentage"`
	FunctionalNp                 null.Float64 `name:"functional_np" bson:"functional_np,omitempty" json:"functional_np"`

	// "Lab Calculation Sheet"
	H2so4Ap                null.Float64 `name:"h2so4_ap" bson:"h2so4_ap,omitempty" json:"h2so4_ap"`
	H2so4Bt                null.Float64 `name:"h2so4_bt" bson:"h2so4_bt,omitempty" json:"h2so4_bt"`
	NAsam                  null.Float64 `name:"n_asam" bson:"n_asam,omitempty" json:"n_asam"`
	SampleVolumeAlkalinity null.Float64 `name:"sample_volume_alkalinity" bson:"sample_volume_alkalinity,omitempty" json:"sample_volume_alkalinity"`
	PtAlkalinityDiff       null.Float64 `name:"pt_alkalinity_diff" bson:"pt_alkalinity_diff,omitempty" json:"pt_alkalinity_diff"`
	PAlkalinity            null.Float64 `name:"p_alkalinity" bson:"p_alkalinity,omitempty" json:"p_alkalinity"`
	TAlkalinity            null.Float64 `name:"t_alkalinity" bson:"t_alkalinity,omitempty" json:"t_alkalinity"`
	NitritDilutionFactor   null.Float64 `name:"nitrit_dilution_factor" bson:"nitrit_dilution_factor,omitempty" json:"nitrit_dilution_factor"`
	NitritBlanko           null.Float64 `name:"nitrit_blanko" bson:"nitrit_blanko,omitempty" json:"nitrit_blanko"`
	NitritAbsAs            null.Float64 `name:"nitrit_abs_as" bson:"nitrit_abs_as,omitempty" json:"nitrit_abs_as"`
	NitritSampleVolume     null.Float64 `name:"nitrit_sample_volume" bson:"nitrit_sample_volume,omitempty" json:"nitrit_sample_volume"`
	NitritConcC            null.Float64 `name:"nitrit_conc_c" bson:"nitrit_conc_c,omitempty" json:"nitrit_conc_c"`
	NitritDescription      null.Float64 `name:"nitrit_description" bson:"nitrit_description,omitempty" json:"nitrit_description"`
	FosfatDilutionFactor   null.Float64 `name:"fosfat_dilution_factor" bson:"fosfat_dilution_factor,omitempty" json:"fosfat_dilution_factor"`
	FosfatBlanko           null.Float64 `name:"fosfat_blanko" bson:"fosfat_blanko,omitempty" json:"fosfat_blanko"`
	FosfatAbsAs            null.Float64 `name:"fosfat_abs_as" bson:"fosfat_abs_as,omitempty" json:"fosfat_abs_as"`
	FosfatSampleVolume     null.Float64 `name:"fosfat_sample_volume" bson:"fosfat_sample_volume,omitempty" json:"fosfat_sample_volume"`
	FosfatConcC            null.Float64 `name:"fosfat_conc_c" bson:"fosfat_conc_c,omitempty" json:"fosfat_conc_c"`
	FosfatDescription      null.Float64 `name:"fosfat_description" bson:"fosfat_description,omitempty" json:"fosfat_description"`
	TanDilutionFactor      null.Float64 `name:"tan_dilution_factor" bson:"tan_dilution_factor,omitempty" json:"tan_dilution_factor"`
	TanBlanko              null.Float64 `name:"tan_blanko" bson:"tan_blanko,omitempty" json:"tan_blanko"`
	TanAbsAs               null.Float64 `name:"tan_abs_as" bson:"tan_abs_as,omitempty" json:"tan_abs_as"`
	TanSampleVolume        null.Float64 `name:"tan_sample_volume" bson:"tan_sample_volume,omitempty" json:"tan_sample_volume"`
	TanConcC               null.Float64 `name:"tan_conc_c" bson:"tan_conc_c,omitempty" json:"tan_conc_c"`
	TanDescription         null.Float64 `name:"tan_description" bson:"tan_description,omitempty" json:"tan_description"`
	UnionizedTan           null.Float64 `name:"unionized_tan" bson:"unionized_tan,omitempty" json:"unionized_tan"`
	I                      null.Float64 `name:"i" bson:"i,omitempty" json:"i"`
	PkaS                   null.Float64 `name:"pka_s" bson:"pka_s,omitempty" json:"pka_s"`
	FractionNh3            null.Float64 `name:"fraction_nh3" bson:"fraction_nh3,omitempty" json:"fraction_nh3"`
	FractionNh4            null.Float64 `name:"fraction_nh4" bson:"fraction_nh4,omitempty" json:"fraction_nh4"`
	XVolume                null.Float64 `name:"x_volume" bson:"x_volume,omitempty" json:"x_volume"`
	YVolume                null.Float64 `name:"y_volume" bson:"y_volume,omitempty" json:"y_volume"`
	NKmno4                 null.Float64 `name:"n_kmno4" bson:"n_kmno4,omitempty" json:"n_kmno4"`
	BEKmn04                null.Float64 `name:"be_kmn04" bson:"be_kmn04,omitempty" json:"be_kmn04"`
	SampleVolumeTitration  null.Float64 `name:"sample_volume_titration" bson:"sample_volume_titration,omitempty" json:"sample_volume_titration"`
	CuVolume               null.Float64 `name:"cu_volume" bson:"cu_volume,omitempty" json:"cu_volume"`
	VEdtaA                 null.Float64 `name:"v_edta_a" bson:"v_edta_a,omitempty" json:"v_edta_a"`
	VEdtaB                 null.Float64 `name:"v_edta_b" bson:"v_edta_b,omitempty" json:"v_edta_b"`
	MEdta                  null.Float64 `name:"m_edta" bson:"m_edta,omitempty" json:"m_edta"`
	MrCaCO3                null.Float64 `name:"mr_caco3" bson:"mr_caco3,omitempty" json:"mr_caco3"`
	ArCa                   null.Float64 `name:"ar_ca" bson:"ar_ca,omitempty" json:"ar_ca"`
	ArMg                   null.Float64 `name:"ar_mg" bson:"ar_mg,omitempty" json:"ar_mg"`

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
	OtherGreenAlgae     null.Float64 `name:"other_green_algae" bson:"other_green_algae,omitempty" json:"other_green_algae"`
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
	OtherBlueGreenAlgae null.Float64 `name:"other_blue_green_algae" bson:"other_blue_green_algae,omitempty" json:"other_blue_green_algae"`
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
	OtherChrisophyta    null.Float64 `name:"other_chrisophyta" bson:"other_chrisophyta,omitempty" json:"other_chrisophyta"`
	Astasia             null.Float64 `name:"astasia" bson:"astasia,omitempty" json:"astasia"`
	Cryptoglena         null.Float64 `name:"cryptoglena" bson:"cryptoglena,omitempty" json:"cryptoglena"`
	Euglena             null.Float64 `name:"euglena" bson:"euglena,omitempty" json:"euglena"`
	Trachelomonas       null.Float64 `name:"trachelomonas" bson:"trachelomonas,omitempty" json:"trachelomonas"`
	Phacus              null.Float64 `name:"phacus" bson:"phacus,omitempty" json:"phacus"`
	Rynchomonas         null.Float64 `name:"rynchomonas" bson:"rynchomonas,omitempty" json:"rynchomonas"`
	OtherEuglenophyta   null.Float64 `name:"other_euglenophyta" bson:"other_euglenophyta,omitempty" json:"other_euglenophyta"`
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
	OtherDinoflagellata null.Float64 `name:"other_dinoflagellata" bson:"other_dinoflagellata,omitempty" json:"other_dinoflagellata"`
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
	OtherZooplankton    null.Float64 `name:"other_zooplankton" bson:"other_zooplankton,omitempty" json:"other_zooplankton"`
	Cryptomonas         null.Float64 `name:"cryptomonas" bson:"cryptomonas,omitempty" json:"cryptomonas"`
	Prymenesium         null.Float64 `name:"prymenesium" bson:"prymenesium,omitempty" json:"prymenesium"`
	Isochrysis          null.Float64 `name:"isochrysis" bson:"isochrysis,omitempty" json:"isochrysis"`
	OtherPlankton       null.Float64 `name:"other_plankton" bson:"other_plankton,omitempty" json:"other_plankton"`
	BacteryTotal        null.Float64 `name:"bactery_total" bson:"bactery_total,omitempty" json:"bactery_total"`

	// Bactery TVC and TBC (In Lab Calculation Sheet)
	InoculumTvc1            null.Float64 `name:"inoculum_tvc_1" bson:"inoculumTvc1,omitempty" json:"inoculum_tvc_1"`
	DilutionFactorTvc1      null.Float64 `name:"dilution_factor_tvc_1" bson:"dilutionFactorTvc1,omitempty" json:"dilution_factor_tvc_1"`
	YellowPetri1Tvc1        null.Float64 `name:"yellow_petri1_tvc_1" bson:"yellowPetri1Tvc1,omitempty" json:"yellow_petri1_tvc_1"`
	GreenPetri1Tvc1         null.Float64 `name:"green_petri1_tvc_1" bson:"greenPetri1Tvc1,omitempty" json:"green_petri1_tvc_1"`
	BlackPetri1Tvc1         null.Float64 `name:"black_petri1_tvc_1" bson:"blackPetri1Tvc1,omitempty" json:"black_petri1_tvc_1"`
	AdjustedCountPetri1Tvc1 null.Float64 `name:"adjusted_count_petri1_tvc_1" bson:"adjustedCountPetri1Tvc1,omitempty" json:"adjusted_count_petri1_tvc_1"`
	YellowPetri2Tvc1        null.Float64 `name:"yellow_petri2_tvc_1" bson:"yellowPetri2Tvc1,omitempty" json:"yellow_petri2_tvc_1"`
	GreenPetri2Tvc1         null.Float64 `name:"green_petri2_tvc_1" bson:"greenPetri2Tvc1,omitempty" json:"green_petri2_tvc_1"`
	BlackPetri2Tvc1         null.Float64 `name:"black_petri2_tvc_1" bson:"blackPetri2Tvc1,omitempty" json:"black_petri2_tvc_1"`
	AdjustedCountPetri2Tvc1 null.Float64 `name:"adjusted_count_petri2_tvc_1" bson:"adjustedCountPetri2Tvc1,omitempty" json:"adjusted_count_petri2_tvc_1"`
	TotalCountTvc1          null.Float64 `name:"total_count_tvc_1" bson:"totalCountTvc1,omitempty" json:"total_count_tvc_1"`
	DilutionFactorTvc2      null.Float64 `name:"dilution_factor_tvc_2" bson:"dilutionFactorTvc2,omitempty" json:"dilution_factor_tvc_2"`
	YellowPetri1Tvc2        null.Float64 `name:"yellow_petri1_tvc_2" bson:"yellowPetri1Tvc2,omitempty" json:"yellow_petri1_tvc_2"`
	GreenPetri1Tvc2         null.Float64 `name:"green_petri1_tvc_2" bson:"greenPetri1Tvc2,omitempty" json:"green_petri1_tvc_2"`
	BlackPetri1Tvc2         null.Float64 `name:"black_petri1_tvc_2" bson:"blackPetri1Tvc2,omitempty" json:"black_petri1_tvc_2"`
	AdjustedCountPetri1Tvc2 null.Float64 `name:"adjusted_count_petri1_tvc_2" bson:"adjustedCountPetri1Tvc2,omitempty" json:"adjusted_count_petri1_tvc_2"`
	TotalCountTvc2          null.Float64 `name:"total_count_tvc_2" bson:"totalCountTvc2,omitempty" json:"total_count_tvc_2"`
	InoculumTbc1            null.Float64 `name:"inoculum_tbc_1" bson:"inoculumTbc1,omitempty" json:"inoculum_tbc_1"`
	DilutionFactorTbc1      null.Float64 `name:"dilution_factor_tbc_1" bson:"dilutionFactorTbc1,omitempty" json:"dilution_factor_tbc_1"`
	Petri1Tbc1              null.Float64 `name:"petri1_tbc_1" bson:"petri1Tbc1,omitempty" json:"petri1_tbc_1"`
	AdjustedCountPetri1Tbc1 null.Float64 `name:"adjusted_count_petri1_tbc_1" bson:"adjustedCountPetri1Tbc1,omitempty" json:"adjusted_count_petri1_tbc_1"`
	Petri2Tbc1              null.Float64 `name:"petri2_tbc_1" bson:"petri2Tbc1,omitempty" json:"petri2_tbc_1"`
	AdjustedCountPetri2Tbc1 null.Float64 `name:"adjusted_count_petri2_tbc_1" bson:"adjustedCountPetri2Tbc1,omitempty" json:"adjusted_count_petri2_tbc_1"`
	TotalCountTbc1          null.Float64 `name:"total_count_tbc_1" bson:"totalCountTbc1,omitempty" json:"total_count_tbc_1"`
	DilutionFactorTbc2      null.Float64 `name:"dilution_factor_tbc_2" bson:"dilutionFactorTbc2,omitempty" json:"dilution_factor_tbc_2"`
	Petri1Tbc2              null.Float64 `name:"petri1_tbc_2" bson:"petri1Tbc2,omitempty" json:"petri1_tbc_2"`
	AdjustedCountPetri1Tbc2 null.Float64 `name:"adjusted_count_petri1_tbc_2" bson:"adjustedCountPetri1Tbc2,omitempty" json:"adjusted_count_petri1_tbc_2"`
	Petri2Tbc2              null.Float64 `name:"petri2_tbc_2" bson:"petri2Tbc2,omitempty" json:"petri2_tbc_2"`
	AdjustedCountPetri2Tbc2 null.Float64 `name:"adjusted_count_petri2_tbc_2" bson:"adjustedCountPetri2Tbc2,omitempty" json:"adjusted_count_petri2_tbc_2"`
	TotalCountTbc2          null.Float64 `name:"total_count_tbc_2" bson:"totalCountTbc2,omitempty" json:"total_count_tbc_2"`

	// Calibration Curvature
	NitritSlope      null.Float64 `name:"nitrit_slope" bson:"nitrit_slope,omitempty" json:"nitrit_slope"`
	NitritIntercept  null.Float64 `name:"nitrit_intercept" bson:"nitrit_intercept,omitempty" json:"nitrit_intercept"`
	PhospatSlope     null.Float64 `name:"phospat_slope" bson:"phospat_slope,omitempty" json:"phospat_slope"`
	PhospatIntercept null.Float64 `name:"phospat_intercept" bson:"phospat_intercept,omitempty" json:"phospat_intercept"`
	TANSlope         null.Float64 `name:"tan_slope" bson:"tan_slope,omitempty" json:"tan_slope"`
	TANIntercept     null.Float64 `name:"tan_intercept" bson:"tan_intercept,omitempty" json:"tan_intercept"`

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

type AggridMetricInput struct {
	ID            string      `json:"id"`
	AttributeName string      `json:"attribute_name"`
	Value         interface{} `json:"value"`
}
