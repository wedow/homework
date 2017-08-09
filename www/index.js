document.body.onload = function() {
	fetch('/text').then(resp => resp.text()).then(resp => {
		let posts = JSON.parse(resp)
		if (posts)
			posts.forEach(post => renderPost(post))
	})
}

document.querySelector('.editor').onsubmit = submitEditor

function submitEditor(e, parent_id) {
	e.preventDefault()

	let body = e.currentTarget.querySelector('form [name=text-input]')
	let user = e.currentTarget.querySelector('form [name=username]')
	let city = e.currentTarget.querySelector('form [name=city]')

	if (!user.value || !body.value || !city.value)
		return alert('Please fill in the User, City, and Content fields')

	let data = {
		username: user.value,
		content: body.value,
		city: city.value
	}

	if (parent_id)
		data.parent_id = parent_id

	fetch('/text',{
		headers: new Headers({'Content-Type': 'application/json'}),
		method: 'post',
		body: JSON.stringify(data)
	}).then(resp => resp.text()).then(resp => {
		let post = JSON.parse(resp)
		renderPost(post)
		body.value = ''
	})
}

var postStore = {}

function renderPost(post) {
	let div = document.createElement('div')
	let metadata = document.createElement('div')
	let content = document.createElement('div')
	let hr = document.createElement('hr')

	metadata.append(' - ', post.username, createReplyLink(post.id))
	content.append(post.content)
	content.className = 'content'

	div.append(content, metadata, hr)
	div.className = 'post'

	postStore[post.id] = div
	if (post.parent_id) {
		postStore[post.parent_id].append(div)
	} else {
		document.querySelector('#posts').prepend(div)
	}
}

function createReplyLink(id) {
	let link = document.createElement('a')
	link.append('Reply')
	link.onclick = function(e) {
		e.preventDefault()
		let original = document.querySelector('.editor')
		let editor = original.cloneNode(true)

		let swapBackElements = () => editor.replaceWith(link)

		let cancelButton = document.createElement('button')
		cancelButton.append('Cancel')
		cancelButton.onclick = swapBackElements

		editor.append(cancelButton)
		editor.onsubmit = function(e) {
			submitEditor(e, id)
			swapBackElements()
		}

		link.replaceWith(editor)
		editor.querySelector('[name=username]').focus()
	}

	return link
}
