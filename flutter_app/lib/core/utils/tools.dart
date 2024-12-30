import 'package:flutter/material.dart';
import 'package:squad_go/core/models/team.dart';

Color getColorBasedOnDate(String date) {
  final DateTime eventDate = DateTime.parse(date);
  final DateTime now = DateTime.now();

  if (eventDate.isBefore(now.subtract(const Duration(days: 1)))) {
    return Colors.grey;
  } else if (eventDate.day == now.day &&
      eventDate.month == now.month &&
      eventDate.year == now.year) {
    return Colors.blue;
  } else {
    return Colors.green;
  }
}

bool hasRole(
  PlayerRole role,
  String? userID,
  List<Team>? teams,
) {
  if (teams == null || userID == null) return false;

  for (final team in teams) {
    for (final player in team.players) {
      if (player.userID == userID && player.role == role) {
        return true;
      }
    }
  }

  return false;
}