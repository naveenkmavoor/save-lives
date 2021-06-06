import 'dart:math';
import 'package:flutter/material.dart';
import 'package:flutter/rendering.dart';

class MyAppBar extends StatefulWidget implements PreferredSizeWidget {
  Size get preferredSize => new Size.fromHeight(68);
  @override
  _MyAppBarState createState() => _MyAppBarState();
}

class _MyAppBarState extends State<MyAppBar> {
 

  @override
  Widget build(BuildContext context) {
    final screenSize = MediaQuery.of(context).size;
    return AppBar(
      automaticallyImplyLeading: false,
      toolbarHeight: 68,
      titleSpacing: pow((screenSize.width / 79), 2),
      title: Row(
        children: [
          Text(
            'Save Lives Dashboard',
            style: TextStyle(
              fontSize: 18,
              letterSpacing: 1,
            ),
          ),
          SizedBox(
            width: 5,
          ),
        ],
      ),
    );
  }
}

