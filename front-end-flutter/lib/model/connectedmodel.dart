import 'package:save_lives/class/case.dart';
import 'package:scoped_model/scoped_model.dart';
import 'dart:convert' as convert;

import 'package:http/http.dart' as http;

class ConnectedModel extends Model {
  List<Case> caselist = [];
  bool isloading = true;

  Future<List<Case>> fetchCases() async {
    var url = Uri.parse("http://localhost:8083/case");
    List<dynamic> responseValue = [];
    try {
      http.Response response = await http.get(url);
      if (response.statusCode == 200) {
        responseValue = convert.jsonDecode(response.body);
      }

      caselist =
          responseValue.map<Case>((json) => Case.fromJson(json)).toList();

      return caselist.reversed.toList();
    } catch (e) {
      print(e);
      return caselist;
    }
  }

  Future<Case> fetchOneCase(String id) async {
    Case onecase;
    var url = Uri.parse("http://localhost:8083/case/$id");
    var responseValue = {};
    try {
      http.Response response = await http.get(url);
      if (response.statusCode == 200) {
        responseValue = convert.jsonDecode(response.body);
      }
      print("responsevalue splitting : $responseValue ");
      onecase = Case.fromJson(responseValue);
      print("caseval : ${onecase.accidentLoc}");

      return onecase;
    } catch (e) {
      print(e);
      return onecase;
    }
  }
}
