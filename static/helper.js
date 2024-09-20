function copyToClipboard(text){

	if (navigator.clipboard && navigator.clipboard.writeText) {

		navigator.clipboard.writeText(text)

	}else{

		//sometimes clipboard api has problems on mobile browsers or won't work unless https connection, fallback
		//into this method instead
		let element = document.createElement("textarea");
		element.value = text;
		document.body.appendChild(element);
		element.select();
		document.execCommand("copy");
		document.body.removeChild(element);

	}

}
