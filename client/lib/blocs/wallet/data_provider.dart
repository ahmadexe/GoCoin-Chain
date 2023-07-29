part of 'wallet_bloc.dart';

class _WallterProvider {
  static final _handler = Dio();

  static Future<Wallet> getWalletDetails() async {
    try {
      final response = await _handler.post(
        'http://10.0.2.2:5050/wallet',
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
