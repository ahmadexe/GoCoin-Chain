part of 'transaction_bloc.dart';

class _TransactionProvider {
  static final _handler = Dio();
  static Future<void> createTransaction(Map<String, dynamic> payload) async {
    try {
      await _handler.post(
        'http://0.0.0.0:8080/transaction',
        data: payload,
      );
    } catch (e) {
      debugPrint('----- ERROR in Create Transaction Provider -----');
      debugPrint(e.toString());
      throw Exception('Failed to create transaction');
    }
  }
}
