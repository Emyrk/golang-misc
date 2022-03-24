from coinbase.wallet.client import Client
client = Client("J5AUYanhoXzdUJD6", "X2Qig65SLfBf3bNLYc2U4NlEAAlbg6WC")

notification = client.get_accounts()
print(notification)