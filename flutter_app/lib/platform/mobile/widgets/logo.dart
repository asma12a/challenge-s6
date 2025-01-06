import 'package:flutter/material.dart';

class Logo extends StatelessWidget {
  final double? width;

  const Logo({super.key, this.width});

  @override
  Widget build(BuildContext context) {
    return Image(
      image: AssetImage('assets/images/logo.png'),
      width: width,
    );
  }
}
