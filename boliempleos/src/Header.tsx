import logoSrc from "/logo.svg";

export default function Header() {
	return (
		<header class="py-2">
			<div class="flex justify-center">
				<img src={logoSrc} class="w-24 h-24" />
			</div>
			<h1 class="text-xl text-center uppercase font-normal tracking-tight">Boliempleos</h1>
			<p class="text-center text-sm">Ofertas de empleo de las empresas más grandes de Bolivia</p>
		</header>
	)
}
