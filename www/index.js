document.body.onload = function () {
  window.fetch('/text').then(resp => resp.text()).then(resp => {
    let posts = JSON.parse(resp)
    if (posts) { posts.forEach(renderPost) }
  })
}

document.querySelector('.editor').onsubmit = submitEditor

function submitEditor (e, parentId) {
  e.preventDefault()

  let content = e.currentTarget.querySelector('form [name=text-input]').value
  let username = e.currentTarget.querySelector('form [name=username]').value
  let city = e.currentTarget.querySelector('form [name=city]').value

  if (!username || !content || !city) {
    return window.alert('Please fill in the User, City, and Content fields')
  }

  let data = { username, content, city }

  if (parentId) { data.parent_id = parentId }

  window.fetch('/text', {
    headers: new window.Headers({'Content-Type': 'application/json'}),
    method: 'post',
    body: JSON.stringify(data)
  }).then(resp => resp.text()).then(resp => {
    renderPost(JSON.parse(resp))
    body.value = ''
  })
}

const postStore = {}

function renderPost (post) {
  let timestamp = ' (' + window.moment(post.created_at).fromNow() + ')'
  let location = ' in ' + post.city + ' (lat: ' + post.latitude + ' long: ' + post.longitude + ') Temperature: ' + post.temperature + '°C'
  let metadata = document.createElement('div')
  metadata.append(
    ' ~ ',
    post.username,
    timestamp,
    location,
    createReplyLink(post.id)
  )

  let content = document.createElement('div')
  content.append(post.content)
  content.className = 'content'

  let hr = document.createElement('hr')

  let div = document.createElement('div')
  div.append(content, metadata, hr)
  div.className = 'post'

  postStore[post.id] = div
  if (post.parent_id) {
    postStore[post.parent_id].append(div)
  } else {
    document.querySelector('#posts').prepend(div)
  }
}

function createReplyLink (id) {
  let link = document.createElement('a')
  link.append('Reply')
  link.onclick = (e) => {
    e.preventDefault()
    let editor = document.querySelector('.editor').cloneNode(true)

    let cancelButton = document.createElement('button')
    cancelButton.append('Cancel')
    cancelButton.onclick = () => editor.replaceWith(link)

    editor.append(cancelButton)
    editor.onsubmit = (e) => {
      submitEditor(e, id)
      editor.replaceWith(link)
    }

    link.replaceWith(editor)
    editor.querySelector('[name=username]').focus()
  }

  return link
}
