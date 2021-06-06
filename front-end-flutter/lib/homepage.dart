import 'package:flutter/material.dart';
import 'package:save_lives/class/case.dart';

import 'package:save_lives/model/connectedmodel.dart';
import 'package:save_lives/secondpage.dart';
import 'package:save_lives/widgets/appbar.dart';

class Homepage extends StatefulWidget {
  static const String route = '/home';

  final ConnectedModel model;
  Homepage(this.model);

  @override
  _HomepageState createState() => _HomepageState();
}

class _HomepageState extends State<Homepage> {
  ConnectedModel model;
  @override
  void initState() {
    model = widget.model;
    super.initState();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: MyAppBar(),
        body: SizedBox.expand(
          child: Row(children: [
            Container(
              width: MediaQuery.of(context).size.width * 0.2,
              child: Card(
                child: Column(
                  children: [
                    InkWell(
                      autofocus: true,
                      onTap: () {},
                      borderRadius: BorderRadius.circular(10),
                      child: ListTile(
                        leading: Icon(Icons.report),
                        title: Text("Local Reports"),
                      ),
                    ),
                    ListTile(
                      leading: Icon(Icons.notifications),
                      title: Text("Global Reports"),
                    ),
                  ],
                ),
              ),
            ),
            Container(
              width: MediaQuery.of(context).size.width * 0.7,
              child: Stack(
                children: [
                  FutureBuilder<List<Case>>(
                      future: model.fetchCases(),
                      builder: (BuildContext context, AsyncSnapshot snapshot) {
                        if (snapshot.connectionState ==
                            ConnectionState.waiting) {
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
                          return ListView.builder(
                            itemBuilder: (BuildContext context, int index) {
                              int c = snapshot.data.length - index;
                              return Card(
                                child: ListTile(
                                  onTap: () {
                                    String val = snapshot.data[c - 1].id;
                                    // _model.caselist[index].isPressed = true;
                                    Navigator.of(context).push(
                                        (MaterialPageRoute(
                                            builder: (context) =>
                                                SecondPage(val, model))));
                                  },
                                  title: Text("Case Alert #$c"),
                                  trailing: Text(
                                      'case ID: ${snapshot.data[c - 1].id}'),
                                ),
                              );
                            },
                            itemCount: snapshot.data.length,
                          );
                        }
                        return Container();
                      })
                ],
              ),
            ),
          ]),
        ));
  }
}
