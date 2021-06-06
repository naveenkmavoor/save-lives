import 'package:flutter/material.dart';
import 'package:save_lives/homepage.dart';
import 'package:save_lives/model/connectedmodel.dart';
import 'package:save_lives/theme/styles.dart';
import 'package:scoped_model/scoped_model.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  final ConnectedModel _model = ConnectedModel();
  @override
  Widget build(BuildContext context) {
    return ScopedModel(
      model: _model,
      child: MaterialApp(debugShowCheckedModeBanner: false,
        theme: Styles.themeData(context),
        home: Scaffold(
          body: Homepage(_model),
        ),
      ),
    );
  }
}
