import 'package:flutter/material.dart';
import 'package:squad_go/core/models/team.dart';

Color getColorBasedOnDate(String date) {
  final DateTime now = DateTime.now();
  final DateTime eventDate = DateTime.parse(date).toLocal(); // Assure que la date est bien locale

  if (eventDate.isBefore(DateTime(now.year, now.month, now.day))) {
    return Colors.grey; // Événement passé (hier ou avant)
  } else if (eventDate.year == now.year &&
      eventDate.month == now.month &&
      eventDate.day == now.day) {
    // L'événement est aujourd'hui
    if (eventDate.isAfter(now)) {
      return Colors.green; // Pas encore commencé
    } else {
      return Colors.blue; // Déjà commencé
    }
  } else {
    return Colors.green; // Événement prévu pour demain ou plus tard
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
