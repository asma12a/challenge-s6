abstract class Constants {
  static const String apiBaseUrl = String.fromEnvironment(
    'API_BASE_URL',
    defaultValue: 'http://127.0.0.1:3001',
  );

  static const String jwtStorageToken = String.fromEnvironment(
    'JWT_STORAGE_KEY',
    defaultValue: 'squadgo-jwt',
  );
}
