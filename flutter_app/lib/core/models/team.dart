class Team {
  final String id;
  final String name;
  final int maxPlayers;
  final List<Player> players;

  const Team({
    required this.id,
    required this.name,
    required this.maxPlayers,
    this.players = const [],
  });

  factory Team.fromJson(Map<String, dynamic> data) {
    return Team(
      id: data['id'],
      name: data['name'],
      maxPlayers: data['maxPlayers'] ?? 0,
      players: data['players'] != null
          ? (data['players'] as List<dynamic>)
              .map((player) => Player.fromJson(player))
              .toList()
          : [],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'maxPlayers': maxPlayers,
      'players': players.map((player) => player.toJson()).toList(),
    };
  }
}

class Player {
  final String id;
  final String? name;
  final String email;
  final PlayerRole role;
  final PlayerStatus status;
  final String? userID;

  const Player({
    required this.id,
    required this.email,
    required this.role,
    required this.status,
    this.name,
    this.userID,
  });

  factory Player.fromJson(Map<String, dynamic> data) {
    return Player(
      id: data['id'],
      name: data['name'],
      email: data['email'],
      role: PlayerRole.values.firstWhere((e) =>
          e.toString().split('.').last.toLowerCase() ==
          data['role'].toLowerCase()),
      status: PlayerStatus.values.firstWhere((e) =>
          e.toString().split('.').last.toLowerCase() ==
          data['status'].toLowerCase()),
      userID: data['user_id'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'email': email,
      'role': role.toString().split('.').last,
      'status': status.toString().split('.').last,
      'user_id': userID,
    };
  }
}

enum PlayerRole { player, coach, org }

enum PlayerStatus { pending, valid }
