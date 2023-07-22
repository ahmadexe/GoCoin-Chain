part of 'wallet_bloc.dart';

class _WallterProvider {
  final handler = Dio();

  static Future<Wallet> getWalletDetails() async {
    try {
      final response = await Dio().post(
        'http://0.0.0.0:8080/wallet',
      );

      final wallet = Wallet.fromMap(response.data as Map<String, dynamic>);

      return wallet;
    } catch (e) {
      debugPrint('----- ERROR in Wallet Details Provider -----');
      debugPrint(e.toString());
      throw Exception('Failed to get wallet details');
    }
  }
}
