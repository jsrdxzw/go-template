// Code generated by ent, DO NOT EDIT.

package entschema

import (
	"example/be/gen/entschema/photographcity"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// PhotographCity is the model entity for the PhotographCity schema.
type PhotographCity struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// 省份id
	ProvinceCode *string `json:"province_code,omitempty"`
	// 省份名称
	ProvinceName *string `json:"province_name,omitempty"`
	// 城市id
	CityCode string `json:"city_code,omitempty"`
	// 城市名称
	CityName     string `json:"city_name,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PhotographCity) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case photographcity.FieldID:
			values[i] = new(sql.NullInt64)
		case photographcity.FieldProvinceCode, photographcity.FieldProvinceName, photographcity.FieldCityCode, photographcity.FieldCityName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PhotographCity fields.
func (pc *PhotographCity) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case photographcity.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pc.ID = int(value.Int64)
		case photographcity.FieldProvinceCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field province_code", values[i])
			} else if value.Valid {
				pc.ProvinceCode = new(string)
				*pc.ProvinceCode = value.String
			}
		case photographcity.FieldProvinceName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field province_name", values[i])
			} else if value.Valid {
				pc.ProvinceName = new(string)
				*pc.ProvinceName = value.String
			}
		case photographcity.FieldCityCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city_code", values[i])
			} else if value.Valid {
				pc.CityCode = value.String
			}
		case photographcity.FieldCityName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city_name", values[i])
			} else if value.Valid {
				pc.CityName = value.String
			}
		default:
			pc.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the PhotographCity.
// This includes values selected through modifiers, order, etc.
func (pc *PhotographCity) Value(name string) (ent.Value, error) {
	return pc.selectValues.Get(name)
}

// Update returns a builder for updating this PhotographCity.
// Note that you need to call PhotographCity.Unwrap() before calling this method if this PhotographCity
// was returned from a transaction, and the transaction was committed or rolled back.
func (pc *PhotographCity) Update() *PhotographCityUpdateOne {
	return NewPhotographCityClient(pc.config).UpdateOne(pc)
}

// Unwrap unwraps the PhotographCity entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pc *PhotographCity) Unwrap() *PhotographCity {
	_tx, ok := pc.config.driver.(*txDriver)
	if !ok {
		panic("entschema: PhotographCity is not a transactional entity")
	}
	pc.config.driver = _tx.drv
	return pc
}

// String implements the fmt.Stringer.
func (pc *PhotographCity) String() string {
	var builder strings.Builder
	builder.WriteString("PhotographCity(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pc.ID))
	if v := pc.ProvinceCode; v != nil {
		builder.WriteString("province_code=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := pc.ProvinceName; v != nil {
		builder.WriteString("province_name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("city_code=")
	builder.WriteString(pc.CityCode)
	builder.WriteString(", ")
	builder.WriteString("city_name=")
	builder.WriteString(pc.CityName)
	builder.WriteByte(')')
	return builder.String()
}

// PhotographCities is a parsable slice of PhotographCity.
type PhotographCities []*PhotographCity