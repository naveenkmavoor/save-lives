import 'dart:ui';

import 'package:flutter/material.dart';

class Styles {
  static ThemeData themeData(BuildContext context) {
    final lightTextTheme = Theme.of(context).textTheme.apply(
          fontFamily: 'Open Sans',
          bodyColor: Color(0xffa9a9b3),
          displayColor: Color(0xffa9a9b3),
        );

    return ThemeData(
      iconTheme: IconThemeData(color: Color(0xffa9a9b3)),
      textTheme: lightTextTheme,
      canvasColor: Color(0xff292a2d),
      accentColor: Color(0xffF28E13),
      brightness: Brightness.dark,
      buttonTheme: Theme.of(context)
          .buttonTheme
          .copyWith(colorScheme: ColorScheme.dark()),
      appBarTheme: AppBarTheme(
        iconTheme: IconThemeData(color: Color(0xffa9a9b3)),
        textTheme: lightTextTheme,
        backgroundColor: Color(0xff252627),
        elevation: 0.0,
      ),
    );
  }
}
