# Using Marp for Presentations

This guide explains how to use Marp to create and export presentations from Markdown files.

## Prerequisites

1. VS Code installed
2. Marp for VS Code extension installed

## How to View Presentations

1. Open your Markdown file (e.g., `Introduction_to_Go.marp.md`)
2. Click the "Open Preview to the Side" button in VS Code (or press Ctrl+K V)
3. You'll see the presentation preview in slideshow format

## How to Export to Different Formats

1. Open the Command Palette (Ctrl+Shift+P or Cmd+Shift+P)
2. Type "Marp: Export"
3. Choose from available export formats:
   - PDF (best for sharing)
   - PPTX (PowerPoint format)
   - HTML (web presentation)

## Presenting Your Slides

1. Open the Command Palette
2. Type "Marp: Start Present Slide Deck"
3. Use arrow keys to navigate slides
4. Press 'F' for fullscreen
5. Press 'Esc' to exit

## Marp Markdown Tips

- Use `---` to create a new slide
- Use `<!-- _class: lead -->` for title slides
- Use `![bg right:40%](image.jpg)` for background images
- Use `# ` for slide titles
- Use `## ` for subtitles
- Regular Markdown syntax for:
  - Lists
  - Code blocks
  - Tables
  - etc.

## Themes

The current presentation uses the default theme, but you can change it by modifying the front matter:

```yaml
---
marp: true
theme: default  # or 'gaia' or 'uncover'
paginate: true
---
```

## Custom Styling

You can add custom CSS in the front matter:

```yaml
style: |
  .columns {
    display: grid;
    grid-template-columns: repeat(2, minmax(0, 1fr));
    gap: 1rem;
  }
```

## Tips for Better Presentations

1. Keep slides simple and clean
2. Use images when possible
3. Limit text per slide
4. Use bullet points for key concepts
5. Include code examples sparingly

## Resources

- [Marp Documentation](https://marpit.marp.app/)
- [Marp CLI](https://github.com/marp-team/marp-cli)
- [VS Code Marp Extension](https://marketplace.visualstudio.com/items?itemName=marp-team.marp-vscode)
