import 'package:flutter/material.dart';

class Sport {
  final String id;
  final SportName name;
  final SportType type;
  final Color? color;
  final String? imageUrl;
  final int maxTeams;

  const Sport({
    required this.id,
    required this.name,
    required this.maxTeams,
    required this.type,
    this.color,
    this.imageUrl,
  });

  factory Sport.fromJson(Map<String, dynamic> data) {
    return Sport(
      id: data['id'],
      name: SportName.values.firstWhere((e) =>
          e.toString().split('.').last.toLowerCase() ==
          data['name'].toLowerCase()),
      type: SportType.values.firstWhere((e) =>
          e.toString().split('.').last.toLowerCase() ==
          data['type'].toLowerCase()),
      color: data['color'] != null
          ? Color(int.parse('FF${data['color']}', radix: 16))
          : null,
      imageUrl: data['image_url'],
      maxTeams: data['max_teams'] ?? 0,
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name.toString().split('.').last,
      'type': type.toString().split('.').last,
      'color': color?.value.toRadixString(16).substring(2),
      'image_url': imageUrl,
      'max_teams': maxTeams,
    };
  }

  static Sport empty() {
    return const Sport(
      id: '',
      name: SportName.football,
      maxTeams: 0,
      type: SportType.team,
    );
  }
}

enum SportType { individual, team }

enum SportName { football, basketball, tennis, running }

final Map<SportName, String> sportLabel = {
  SportName.football: 'Football',
  SportName.basketball: 'Basketball',
  SportName.tennis: 'Tennis',
  SportName.running: 'Running',
};

final Map<SportName, IconData> sportIcon = {
  SportName.football: Icons.sports_soccer,
  SportName.basketball: Icons.sports_basketball,
  SportName.tennis: Icons.sports_tennis,
  SportName.running: Icons.directions_run,
};
