package entity2

type Method struct {
}

type Driver struct {
}

type MetadataUsecase struct {
  // UsecaseName
  // InportRequestField
  // InportResponseField
  // OutportMethods []Method
  // Transaction
}

type MetadataController struct {
  // Driver
  // Usecases
}

type MetadataGateway struct {
  // Drivers []Driver
  // Methods []Method
}

type MetadataTest struct {
  // Methods []Method
}




