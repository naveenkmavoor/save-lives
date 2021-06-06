class Case {
  String id;
  String status;
  String dateTime;
  int numbers;
  String vehicleId;
  Hospital hospital;
  List<dynamic> accidentLoc;
  // bool isPressed; //if the user is pressed the message or not

  Case({
    this.id,
    this.status,
    this.dateTime,
    this.numbers,
    this.vehicleId,
    this.hospital,
    this.accidentLoc,
  });

  factory Case.fromJson(Map<String, dynamic> responsevalue) {
    return Case(
        hospital: Hospital.fromJson(responsevalue['hospital']),
        id: responsevalue['id'] as String,
        status: responsevalue['status'],
        dateTime: responsevalue['datetime'],
        numbers: responsevalue['numbers'],
        vehicleId: responsevalue['vehicle_id'],
        accidentLoc: responsevalue['accidentlocation']);
  }
}

class Hospital {
  String name;
  List<dynamic> coordinates;
  String phNo;

  Hospital({this.name, this.coordinates, this.phNo});

  factory Hospital.fromJson(Map<String, dynamic> value) {
    return Hospital(
        name: value['name'] as String,
        coordinates: value['location']['coordinates'] as List<dynamic>,
        phNo: value['number'] as String);
  }
}
