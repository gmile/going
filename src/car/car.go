package car

type Car struct {
  DealerID                      string
  Type                          string
  Stock                         string
  VIN                           string
  Year                          string
  Make                          string
  Model                         string
  ModelNumber                   string
  Body                          string
  Trim                          string
  Doors                         string
  ExteriorColor                 string
  InteriorColor                 string
  EngineCylinders               string
  EngineDisplacement            string
  Transmission                  string
  Miles                         string
  SellingPrice                  string
  MSRP                          string
  Certified                     string
  BookValue                     string
  CategorizedOptions            string
  Invoice                       string
  CityMPG                       string
  HighwayMPG                    string
  DateInStock                   string
  Internet_Price                string
  Description                   string
  Options                       string
  Style_Description             string
  Ext_Color_Generic             string
  Ext_Color_Code                string
  Int_Color_Generic             string
  Int_Color_Code                string
  Int_Upholstery                string
  Engine_Block_Type             string
  Engine_Aspiration_Type        string
  Engine_Description            string
  Transmission_Speed            string
  Transmission_Description      string
  Drivetrain                    string
  Fuel_Type                     string
  EPAClassification             string
  Wheelbase_Code                string
  Misc_Price1                   string
  Misc_Price2                   string
  Misc_Price3                   string
  Factory_Codes                 string
  MarketClass                   string
  PassengerCapacity             string
  ExtColorHexCode               string
  IntColorHexCode               string
  EngineDisplacementCubicInches string
  ImageList                     string
}

func (car *Car) Initialize(csv_row []string) {
  car.DealerID                      = csv_row[ 0]
  car.Type                          = csv_row[ 1]
  car.Stock                         = csv_row[ 2]
  car.VIN                           = csv_row[ 3]
  car.Year                          = csv_row[ 4]
  car.Make                          = csv_row[ 5]
  car.Model                         = csv_row[ 6]
  car.ModelNumber                   = csv_row[ 7]
  car.Body                          = csv_row[ 8]
  car.Trim                          = csv_row[ 9]
  car.Doors                         = csv_row[10]
  car.ExteriorColor                 = csv_row[11]
  car.InteriorColor                 = csv_row[12]
  car.EngineCylinders               = csv_row[13]
  car.EngineDisplacement            = csv_row[14]
  car.Transmission                  = csv_row[15]
  car.Miles                         = csv_row[16]
  car.SellingPrice                  = csv_row[17]
  car.MSRP                          = csv_row[18]
  car.Certified                     = csv_row[19]
  car.BookValue                     = csv_row[20]
  car.CategorizedOptions            = csv_row[21]
  car.Invoice                       = csv_row[22]
  car.CityMPG                       = csv_row[23]
  car.HighwayMPG                    = csv_row[24]
  car.DateInStock                   = csv_row[25]
  car.Internet_Price                = csv_row[26]
  car.Description                   = csv_row[27]
  car.Options                       = csv_row[28]
  car.Style_Description             = csv_row[29]
  car.Ext_Color_Generic             = csv_row[30]
  car.Ext_Color_Code                = csv_row[31]
  car.Int_Color_Generic             = csv_row[32]
  car.Int_Color_Code                = csv_row[33]
  car.Int_Upholstery                = csv_row[34]
  car.Engine_Block_Type             = csv_row[35]
  car.Engine_Aspiration_Type        = csv_row[36]
  car.Engine_Description            = csv_row[37]
  car.Transmission_Speed            = csv_row[38]
  car.Transmission_Description      = csv_row[39]
  car.Drivetrain                    = csv_row[40]
  car.Fuel_Type                     = csv_row[41]
  car.EPAClassification             = csv_row[42]
  car.Wheelbase_Code                = csv_row[43]
  car.Misc_Price1                   = csv_row[44]
  car.Misc_Price2                   = csv_row[45]
  car.Misc_Price3                   = csv_row[46]
  car.Factory_Codes                 = csv_row[47]
  car.MarketClass                   = csv_row[48]
  car.PassengerCapacity             = csv_row[49]
  car.ExtColorHexCode               = csv_row[50]
  car.IntColorHexCode               = csv_row[51]
  car.EngineDisplacementCubicInches = csv_row[52]
  car.ImageList                     = csv_row[53]
}
