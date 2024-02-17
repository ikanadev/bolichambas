import Home from "./Home";
import { AppStateProvider } from "./useAppState";

function App() {

	return (
		<AppStateProvider>
			<Home />
		</AppStateProvider>
	)
}

export default App
