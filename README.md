# Pasta Fabric Service (Golang Edition)

A **Golang-based** web service that simulates **Meatball Monsters** working together to produce pasta! 🏭🍝

## Features
- Meatball Monsters generate pasta when they have energy.
- Monsters need rest to regain energy.
- RESTful API with `/produce` and `/rest` endpoints.
- Interactive web interface for easy control.

## Installation & Setup

### 1. Clone the repository
```sh
git clone https://github.com/THIS_REPO
cd pasta-fabric-service
```

### 2. Install Go (if not installed)
Ensure you have Go installed. You can check with:
```sh
go version
```
If not installed, download it from [golang.org](https://golang.org/dl/).

### 3. Run the service
```sh
go run main.go
```
This starts the web server on `http://localhost:8080`

### 4. Open the Web Interface
Open `index.html` in your browser or navigate to:
```
http://localhost:8080
```

## API Endpoints
| Method | Endpoint    | Description                    |
|--------|------------|--------------------------------|
| GET    | `/produce` | Monsters make pasta!          |
| GET    | `/rest`    | Monsters regain energy.       |

## Project Structure
```
/pasta-fabric-service
│── main.go         # Golang web service logic
│── index.html      # Web interface
│── styles.css      # Styling for UI
│── README.md       # Documentation
```

## How to Play
1. Click **"Make Pasta"** to have monsters produce pasta.
2. If they get tired, click **"Rest Monsters"** to recharge them.
3. Enjoy the simulation!

## License
This project is open-source and available under the MIT License.

---

Happy coding! 🚀🍝

