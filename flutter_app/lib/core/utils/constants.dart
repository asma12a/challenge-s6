abstract class Constants {
  static const String apiBaseUrl = String.fromEnvironment(
    'API_BASE_URL',
    defaultValue: 'https://challenge-s6-1.onrender.com',
  );
   static const String apiBaseUrlWs = String.fromEnvironment(
    'API_BASE_URL_WS',
    defaultValue: 'challenge-s6-1.onrender.com',
  );

  static const String jwtStorageToken = String.fromEnvironment(
    'JWT_STORAGE_KEY',
    defaultValue: 'squadgo-jwt',
  );


  
}
