import 'package:flutter/material.dart';
import 'dart:html';

import 'package:google_maps/google_maps.dart';
import 'dart:ui' as ui;

Widget getMap(List<dynamic> coordinate) {
  String htmlId = "7";

  // ignore: undefined_prefixed_name
  ui.platformViewRegistry.registerViewFactory(htmlId, (int viewId) {
    final myLatlng = LatLng(coordinate[1], coordinate[0]);

    final mapOptions = MapOptions()
      ..zoom = 12
      ..center = LatLng(coordinate[1], coordinate[0]);

    final elem = DivElement()
      ..id = htmlId
      ..style.width = "100%"
      ..style.height = "100%"
      ..style.border = 'none' ;

    final map = GMap(elem, mapOptions);

    Marker(MarkerOptions()
      ..position = myLatlng
      ..map = map
      ..title = 'Hello World!');

    return elem;
  });

  return HtmlElementView(viewType: htmlId);
}

class GoogleMaps extends StatelessWidget {
  final List<dynamic> coordinates;
  GoogleMaps(this.coordinates);
  @override
  Widget build(BuildContext context) {
    return getMap(coordinates);
  }
}
