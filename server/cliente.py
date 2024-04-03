import socket as skt

serverAddr = "localhost"
serverPort = 63420

clientSocket = skt.socket(skt.AF_INET, skt.SOCK_DGRAM)

toSend = input("Texto a enviar: ")
if toSend == "STOP":
  clientSocket.sendto(toSend.encode(), (serverAddr, serverPort))
  print("Programa finalizado")
else:
  clientSocket.sendto(toSend.encode(), (serverAddr, serverPort))
  msg, addr = clientSocket.recvfrom(1024)
  print(msg.decode())