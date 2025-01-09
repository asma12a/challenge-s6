import 'package:dio/dio.dart';
import 'package:dio_cache_interceptor/dio_cache_interceptor.dart';
import 'package:flutter_secure_storage/flutter_secure_storage.dart';
import 'package:squad_go/core/exceptions/app_exception.dart';
import 'package:squad_go/core/models/team.dart';
import 'package:squad_go/core/utils/connectivity_handler.dart';
import 'package:squad_go/main.dart';

const apiBaseUrl = String.fromEnvironment('API_BASE_URL');
const jwtStorageToken = String.fromEnvironment('JWT_STORAGE_KEY');

class TeamService {
  final storage = const FlutterSecureStorage();

  final refreshCacheOptions = CacheOptions(
    store: MemCacheStore(),
    policy: CachePolicy.refresh,
  );

  Future<void> createTeam(String eventID, String name, int? maxPlayers) async {
    final token = await storage.read(key: jwtStorageToken);

    try {
      final url = '$apiBaseUrl/api/events/$eventID/teams';

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
    final token = await storage.read(key: jwtStorageToken);

    try {
      final url = '$apiBaseUrl/api/events/$eventID/teams/${team.id}';

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
    final token = await storage.read(key: jwtStorageToken);

    try {
      final url = '$apiBaseUrl/api/events/$eventID/teams/$teamID/join';

      await dio.post(url.toString(),
          options: ConnectivityHandler().isConnected
              ? refreshCacheOptions.toOptions().copyWith(
                  headers: {
                    'Content-Type': 'application/json',
                    'Authorization': "Bearer $token",
                  },
                )
              : Options(
                  headers: {
                    'Content-Type': 'application/json',
                    'Authorization': "Bearer $token",
                  },
                ));
    } catch (error) {
      throw AppException(message: 'Failed to join team, please try again.');
    }
  }

  Future<void> switchTeam(String eventID, String teamID) async {
    final token = await storage.read(key: jwtStorageToken);

    try {
      final url = '$apiBaseUrl/api/events/$eventID/teams/$teamID/switch';

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
    final token = await storage.read(key: jwtStorageToken);
    try {
      final url = '$apiBaseUrl/api/events/$eventID/teams/$teamID/players';

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
    final token = await storage.read(key: jwtStorageToken);

    try {
      final url = '$apiBaseUrl/api/events/$eventID/teams/players/${player.id}';

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
    final token = await storage.read(key: jwtStorageToken);
    try {
      final url = '$apiBaseUrl/api/events/$eventID/teams/players/$playerID';

      await dio.delete(url.toString(),
          options: ConnectivityHandler().isConnected
              ? refreshCacheOptions.toOptions().copyWith(
                  headers: {
                    'Content-Type': 'application/json',
                    'Authorization': "Bearer $token",
                  },
                )
              : Options(
                  headers: {
                    'Content-Type': 'application/json',
                    'Authorization': "Bearer $token",
                  },
                ));
    } catch (error) {
      throw AppException(message: 'Failed to delete player, please try again.');
    }
  }
}
