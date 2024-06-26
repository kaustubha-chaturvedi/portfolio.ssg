|meta|
title: New Post
header-image: https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSbhwzNrngWxm7FE7Q__Dd3jEEeuYeyewgxXA&s
description: This is my new post
author: Kaustubha Chaturvedi
tags: new, post
date: Feb 04, 2024
|meta|

# A repo only contains a MarkDown files
> Yea, I know that maybe this is not a very useful repository.

# Markdown formating

# Heading 1
`# Heading 1`

## Heading 2
`## Heading 2`

### Heading 3
`### Heading 3`

#### Heading 4
`#### Heading 4`

##### Heading 5
`##### Heading 5`

###### Heading 6
`###### Heading 6`

---

# Paragraph

text `Inline Code` text		
~~Mistaken text.~~	
*Italics*	
**Bold**	

```markdown
text `Inline Code` text		
~~Mistaken text.~~	
*Italics*	
**Bold**	
```

---

# Tasks

- [ ] a task list item
- [ ] list syntax required
- [ ] normal **formatting**
- [ ] incomplete
- [x] completed

```markdown
- [ ] a task list item
- [ ] list syntax required
- [ ] normal **formatting**
- [ ] incomplete
- [x] completed
```

---

# Code Blocks

#### Note:
    4 space indention
    makes full-width
    standard code blocks
	
```markdown
    4 space indention
    makes full-width
    standard code blocks
```


```js
document.innerHTML = "I'm js code";

/* I'm comment */
```

```css
.css {
font-size: 25px;
width: 100%;
padding: 5px;
height: 50%;
color: #333;
}

/* I'm comment */
```

---

# List

* List item one
* List item two
    * A nested item


```markdown
* List item one
* List item two
    * A nested item
```

---

# Numbered List

1. Number list item one		
2. Number list item two
3. Number list item three

```markdown
1. Number list item one		
2. Number list item two
3. Number list item three
```

---

# Quote

> Quote
> 
> Second line Quote

```markdown
> Quote
> 
> Second line Quote
```


---

Standard link =  https://igorkowalczyk.github.io <br>
[Custom Text Link](https://igorkowalczyk.github.io)

```markdown
Standard link =  https://igorkowalczyk.github.io
[Custom Text Link](https://igorkowalczyk.github.io)
```

---

# Image

```markdown
![Image](https://github.com/fluidicon.png)
```
![Image](https://github.com/fluidicon.png)

## Change Size
GFM (Github Flavored MarkDown) don't support change size of img but You can use some HTML "img" tags in your Markdown:
									       
`
<img src="https://github.com/fluidicon.png" alt="Image" width="200" height="200"/>
`

<img src="https://github.com/fluidicon.png" alt="Image" width="200" height="200"/>

---

# Table

| Left-Aligned  | Center Aligned  | Right Aligned |
| :------------ |:---------------:| ------------: |
| col 3 is      | some wordy text | $1600         |
| col 2 is      | centered        |   $12         |
| zebra stripes | are neat        |    $1         |

```
| Left-Aligned  | Center Aligned  | Right Aligned |
| :------------ |:---------------:| -----:|
| col 3 is      | some wordy text | $1600 |
| col 2 is      | centered        |   $12 |
| zebra stripes | are neat        |    $1 |
```

# Youtube Video

![](https://youtu.be/xXQ2F7D6PBY?feature=shared)
