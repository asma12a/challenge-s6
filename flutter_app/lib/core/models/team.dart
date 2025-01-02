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
              .map((player) => Player.fromJson(player, teamId: data['id']))
              .toList()
          : [],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'id': id,
      'name': name,
      'max_players': maxPlayers,
      'players': players.map((player) => player.toJson()).toList(),
    };
  }

  Team copyWith({
    String? id,
    String? name,
    int? maxPlayers,
    List<Player>? players,
  }) {
    return Team(
      id: id ?? this.id,
      name: name ?? this.name,
      maxPlayers: maxPlayers ?? this.maxPlayers,
      players: players ?? this.players,
    );
  }
}

class Player {
  final String id;
  final String? name;
  final String email;
  final PlayerRole role;
  final PlayerStatus status;
  final String? userID;
  final String? teamID;

  const Player({
    required this.id,
    required this.email,
    required this.role,
    required this.status,
    this.name,
    this.userID,
    this.teamID,
  });

  factory Player.fromJson(Map<String, dynamic> data, {String? teamId}) {
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
      teamID: teamId,
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

  Player copyWith({
    String? id,
    String? name,
    String? email,
    PlayerRole? role,
    PlayerStatus? status,
    String? userID,
    String? teamID,
  }) {
    return Player(
      id: id ?? this.id,
      name: name ?? this.name,
      email: email ?? this.email,
      role: role ?? this.role,
      status: status ?? this.status,
      userID: userID ?? this.userID,
      teamID: teamID ?? this.teamID,
    );
  }
}

enum PlayerRole { player, coach, org }

final Map<PlayerRole, String> playerRoleLabel = {
  PlayerRole.player: 'Joueur',
  PlayerRole.coach: 'Coach',
  PlayerRole.org: 'Organisateur',
};

enum PlayerStatus { pending, valid }
