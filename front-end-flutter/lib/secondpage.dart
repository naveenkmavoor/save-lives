import 'package:flutter/material.dart';
import 'package:save_lives/class/case.dart';
import 'package:save_lives/model/connectedmodel.dart';
import 'package:save_lives/widgets/%20map.dart';
import 'package:save_lives/widgets/appbar.dart';

class SecondPage extends StatelessWidget {
  final String id;
  final ConnectedModel model;
  final List<Case> caseReport = [];
  SecondPage(this.id, this.model);

  @override
  Widget build(BuildContext context) {
    final screenSize = MediaQuery.of(context).size;

    return Scaffold(
      appBar: MyAppBar(),
      body: Stack(
        children: [
          FutureBuilder<Case>(
              future: model.fetchOneCase(id),
              builder: (BuildContext context, AsyncSnapshot snapshot) {
                if (snapshot.connectionState == ConnectionState.waiting) {
                  return Center(
                    child: CircularProgressIndicator(),
                  );
                }
                if (snapshot.hasError) {
                  return Center(
                    child: Text("${snapshot.error}"),
                  );
                }

                if (snapshot.hasData) {
                  print("snapshot value is : ${snapshot.data.accidentLoc}");
                  return Align(
                    alignment: Alignment.topCenter,
                    child: Container(
                      width: screenSize.width * 0.5,
                      height: screenSize.height * 0.9,
                      child: Card(
                        child: Column(
                          mainAxisAlignment: MainAxisAlignment.spaceEvenly,
                          children: [
                            Text(
                              'Case Report',
                              style: TextStyle(fontSize: 22),
                            ),
                            Divider(
                              thickness: 2,
                            ),
                            Text(
                                'Accident occured on : ${snapshot.data.dateTime}'),
                            Row(
                              mainAxisAlignment: MainAxisAlignment.center,
                              children: [
                                Container(
                                  height: 20,
                                  width: 20,
                                  decoration: new BoxDecoration(
                                    gradient: LinearGradient(colors: [
                                      Color(0xffFF29F05),
                                      Color(0xffF28705),
                                      Color(0xffF29F05)
                                    ]),
                                    shape: BoxShape.circle,
                                  ),
                                ),
                                SizedBox(
                                  width: 10,
                                ),
                                Text("${snapshot.data.status}")
                              ],
                            ),
                            Container(
                              width: screenSize.width * 0.3,
                              height: screenSize.height * 0.3,
                              child: GoogleMaps(snapshot.data.accidentLoc),
                            ),
                            Text(
                                '${snapshot.data.numbers} Victims Admitted to ${snapshot.data.hospital.name}'),
                            Text('Crashed Vehicle Details : '),
                            Text('Vehicle Reg No.${snapshot.data.vehicleId}'),
                          ],
                        ),
                      ),
                    ),
                  );
                }
                return Container();
              })
        ],
      ),
    );
  }
}
