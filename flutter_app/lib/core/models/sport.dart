import 'package:flutter/material.dart';

class Sport {
  final String id;
  final SportName name;
  final SportType type;
  final Color? color;
  final String? imageUrl;
  final int? maxTeams;

  const Sport({
    required this.id,
    required this.name,
    required this.type,
    this.color,
    this.imageUrl,
    this.maxTeams,
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
      color: data['color'] != null ? Color(data['color']) : null,
      imageUrl: data['image_url'],
      maxTeams: data['max_teams'],
    );
  }
}

enum SportType { individual, team }

enum SportName { football, basketball, tennis, running }

final Map<SportName, IconData> sportIcon = {
  SportName.football: Icons.sports_soccer,
  SportName.basketball: Icons.sports_basketball,
  SportName.tennis: Icons.sports_tennis,
  SportName.running: Icons.directions_run,
};
