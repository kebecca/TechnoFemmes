# TechnoFemmes

# Guide de Test pour l'Application Web en Go

Ce guide explique comment tester notre application web développée en Go.

## Prérequis

- Assurez-vous que vous avez Go installé sur votre machine. Sinon, vous pouvez le télécharger et l'installer à partir du [site officiel de Go](https://golang.org/).

- Vous devez avoir une base de données MySQL configurée avec les informations de connexion appropriées. Assurez-vous que le serveur MySQL est en cours d'exécution.

- Assurez-vous que les fichiers HTML (inscription.html, connexion.html, index.html) sont présents dans le répertoire de travail de votre application.

## Instructions

1. Clonez ce dépôt sur votre machine : https://github.com/kebecca/TechnoFemmes.git


2. Accédez au répertoire du projet : TechnoFemmes


3. Ouvrez le fichier `main.go` dans un éditeur de texte et assurez-vous que les informations de connexion MySQL (utilisateur, mot de passe, hôte, port, base de données) sont correctes.

4. Ouvrez un terminal et exécutez la commande suivante pour exécuter l'application : go run .

5. Ouvrez votre navigateur Web et accédez à l'URL suivante pour accéder à l'application : http://localhost:8080

6. Testez les fonctionnalités de l'application, notamment l'inscription et la connexion.

7. Assurez-vous que les fichiers statiques (CSS, images, JavaScript) sont correctement chargés .

8. Arrêtez l'application en appuyant sur `Ctrl+C` dans le terminal où l'application est en cours d'exécution.

