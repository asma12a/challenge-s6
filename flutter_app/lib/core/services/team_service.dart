import 'package:dio/dio.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:squad_go/core/exceptions/app_exception.dart';
import 'package:squad_go/core/models/team.dart';
import 'package:squad_go/main.dart';

class TeamService {
  final storage = const FlutterSecureStorage();

  Future<void> createTeam(String eventID, String name, int? maxPlayers) async {
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);

    final Uri url =
        Uri.http(dotenv.env['API_BASE_URL']!, 'api/events/$eventID/teams');

    try {
      await dio.post(url.toString(),
          data: {
            'name': name,
            'max_players': maxPlayers,
          },
          options: Options(headers: {
            'Content-Type': 'application/json',
            'Authorization': "Bearer $token",
          }));
    } catch (error) {
      if (error is DioException && error.response != null) {
        throw AppException(
            message: error.response?.data['error'] ??
                'Failed to add player to team, please try again.');
      } else {
        throw AppException(
            message: 'Failed to add player to team, please try again.');
      }
    }
  }

  Future<void> updateTeam(String eventID, Team team) async {
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);

    final Uri url = Uri.http(
        dotenv.env['API_BASE_URL']!, 'api/events/$eventID/teams/${team.id}');

    try {
      await dio.put(url.toString(),
          data: {
            'name': team.name,
            'max_players': team.maxPlayers,
          },
          options: Options(headers: {
            'Content-Type': 'application/json',
            'Authorization': "Bearer $token",
          }));
    } catch (error) {
      throw AppException(message: 'Failed to update team, please try again.');
    }
  }

  Future<void> joinTeam(String eventID, String teamID) async {
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);

    final Uri url = Uri.http(
        dotenv.env['API_BASE_URL']!, 'api/events/$eventID/teams/$teamID/join');

    try {
      await dio.post(url.toString(),
          options: Options(headers: {
            'Content-Type': 'application/json',
            'Authorization': "Bearer $token",
          }));
    } catch (error) {
      throw AppException(message: 'Failed to join team, please try again.');
    }
  }

  Future<void> switchTeam(String eventID, String teamID) async {
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);

    final Uri url = Uri.http(dotenv.env['API_BASE_URL']!,
        'api/events/$eventID/teams/$teamID/switch');

    try {
      await dio.post(url.toString(),
          options: Options(headers: {
            'Content-Type': 'application/json',
            'Authorization': "Bearer $token",
          }));
    } catch (error) {
      throw AppException(message: 'Failed to switch team, please try again.');
    }
  }

  Future<void> addPlayerToTeam(
      String eventID, String teamID, String email, PlayerRole? role) async {
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);

    final Uri url = Uri.http(dotenv.env['API_BASE_URL']!,
        'api/events/$eventID/teams/$teamID/players');

    try {
      await dio.post(url.toString(),
          data: {
            'email': email,
            'role': role?.name,
          },
          options: Options(headers: {
            'Content-Type': 'application/json',
            'Authorization': "Bearer $token",
          }));
    } catch (error) {
      if (error is DioException && error.response != null) {
        // debugPrint(error.response?.data["error"]);
        throw AppException(
            message: error.response?.data['error'] ??
                'Failed to add player to team, please try again.');
      } else {
        throw AppException(
            message: 'Failed to add player to team, please try again.');
      }
    }
  }

  Future<void> updatePlayer(String eventID, Player player) async {
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);

    final Uri url = Uri.http(dotenv.env['API_BASE_URL']!,
        'api/events/$eventID/teams/players/${player.id}');

    try {
      await dio.put(url.toString(),
          data: {
            'role': player.role.name,
            'team_id': player.teamID,
          },
          options: Options(headers: {
            'Content-Type': 'application/json',
            'Authorization': "Bearer $token",
          }));
    } catch (error) {
      throw AppException(message: 'Failed to update player, please try again.');
    }
  }

  Future<void> deletePlayer(String eventID, String playerID) async {
    final token = await storage.read(key: dotenv.env['JWT_STORAGE_KEY']!);
    final Uri url = Uri.http(dotenv.env['API_BASE_URL']!,
        'api/events/$eventID/teams/players/$playerID');
    try {
      await dio.delete(url.toString(),
          options: Options(headers: {
            'Content-Type': 'application/json',
            'Authorization': "Bearer $token",
          }));
    } catch (error) {
      throw AppException(message: 'Failed to delete player, please try again.');
    }
  }
}
