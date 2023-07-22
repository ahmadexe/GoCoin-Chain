part of 'wallet_bloc.dart';

@immutable
abstract class WalletEvent {}

class GetWalletDetails extends WalletEvent {}