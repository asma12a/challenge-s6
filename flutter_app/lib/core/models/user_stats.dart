// import 'package:squad_go/core/models/sport.dart';
import 'package:squad_go/core/models/sport_stat_labels.dart';
import 'package:squad_go/core/models/user.dart';

class UserStats {
  final String? id;
  int value;
  final SportStatLabels? stat;
  final User? user;

  UserStats({this.id, required this.value, this.stat, this.user });

  factory UserStats.fromJson(Map<String, dynamic> json) {
    return UserStats(
      id: json['id'],
      value: json['value'] ?? 0,
      stat: json['stat_label'] != null
          ? SportStatLabels.fromJson(json['stat_label'])
          : null,
      user: json['user'] != null
        ? User.fromJson(json['user'])
          :null
    );
  }

}

class UserPerformance {
  final int nbEvents ;
  final List<UserStats>? stats;

  UserPerformance({
    required this.nbEvents,
    this.stats,
  });

  factory UserPerformance.fromJson(Map<String, dynamic> json) {
    return UserPerformance(
        nbEvents: json['nb_events'],
        stats: json['stats'] != null
            ? List<UserStats>.from(
            json['stats'].map((json) => UserStats.fromJson(json)))
            : null
    );
  }

}
