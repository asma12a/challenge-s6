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
}

class Player {
  final String id;
  final String name;
  final String email;

  const Player({
    required this.id,
    required this.name,
    required this.email,
  });
}
