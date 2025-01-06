import 'dart:ffi';

import 'package:squad_go/core/models/sport.dart';

class SportStatLabels {

  final String? id;
  final String label;
  final String? unit;
  final bool? isMain;

  SportStatLabels({
    this.id,
    required this.label,
    this.unit,
    this.isMain
  });



  factory SportStatLabels.fromJson(Map<String, dynamic> json) {
    return SportStatLabels(
      id: json['id'],
      label: json['label'],
      unit: json['unit'],
      isMain: json['is_main'] ?? false,
    );
  }

}

