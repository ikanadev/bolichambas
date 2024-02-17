import Home from "./Home";
import { DeptoProvider } from "./useDepto";

function App() {

	return (
		<DeptoProvider>
			<Home />
		</DeptoProvider>
	)
}

export default App
