import 'dart:ffi';

import 'package:squad_go/core/models/sport.dart';
import 'package:squad_go/core/models/sport_stat_labels.dart';

class UserStats {

  final String? id;
  final Int value;

  UserStats({
    this.id,
    required this.value
  });



  factory UserStats.fromJson(Map<String, dynamic> json) {
    return UserStats(
      id: json['id'],
      value: json['value']
    );
  }

}

