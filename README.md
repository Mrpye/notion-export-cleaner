# Notion.io cleaner

Simple application to remove the GUID from the filename when exporting [notion.io](http://notion.io/) exports.

I use [notion.io](http://notion.io/) a lot for creating my and managing my documentation but one annoying thing is it appended a GUID to the end of the folders and filenames. This can also cause issues on widows in that it exceeds the max path length if you have complex structured documents.

So hence this tool that simply extracts the files from the zip and removes the GUID from the folders and filename and updated the documents.

 

## Requirements

you will need to install go 1.8 [https://go.dev/doc/install](https://go.dev/doc/install)

## How to install

```yaml
go install https://github.com/Mrpye/notion-export-cleaner
```

## How to run

Export your document from [notion.io](http://notion.io) and clean

1. Click the 3 dots usually top right

![Untitled](Notion%20io%20cleaner/Untitled.png)

1. Click the export

![Untitled](Notion%20io%20cleaner/Untitled%201.png)

1. export the document

![Untitled](Notion%20io%20cleaner/Untitled%202.png)

1. Clean the exported file

```yaml
notion-export-cleaner "./export.zip" "./exported"
```