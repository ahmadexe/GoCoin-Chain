import 'dart:convert';

class Wallet {
  final String publicKey;
  final String privateKey;
  final String blockchainAddress;
  
  Wallet({
    required this.publicKey,
    required this.privateKey,
    required this.blockchainAddress,
  });

  Wallet copyWith({
    String? publicKey,
    String? privateKey,
    String? blockchainAddress,
  }) {
    return Wallet(
      publicKey: publicKey ?? this.publicKey,
      privateKey: privateKey ?? this.privateKey,
      blockchainAddress: blockchainAddress ?? this.blockchainAddress,
    );
  }

  Map<String, dynamic> toMap() {
    return <String, dynamic>{
      'publicKey': publicKey,
      'privateKey': privateKey,
      'blockchainAddress': blockchainAddress,
    };
  }

  factory Wallet.fromMap(Map<String, dynamic> map) {
    return Wallet(
      publicKey: map['publicKey'] as String,
      privateKey: map['privateKey'] as String,
      blockchainAddress: map['blockchainAddress'] as String,
    );
  }

  String toJson() => json.encode(toMap());

  factory Wallet.fromJson(String source) => Wallet.fromMap(json.decode(source) as Map<String, dynamic>);

  @override
  String toString() => 'Wallet(publicKey: $publicKey, privateKey: $privateKey, blockchainAddress: $blockchainAddress)';

  @override
  bool operator ==(covariant Wallet other) {
    if (identical(this, other)) return true;
  
    return 
      other.publicKey == publicKey &&
      other.privateKey == privateKey &&
      other.blockchainAddress == blockchainAddress;
  }

  @override
  int get hashCode => publicKey.hashCode ^ privateKey.hashCode ^ blockchainAddress.hashCode;
}
