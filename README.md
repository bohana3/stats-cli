# stats-cli

'stats-cli' is a tool that performs statistics and aggregation for file metadata provided to it using standard input.
It uses the libary 'stats' that contains the data-structures that allow to perform the operations.

The executable can be run from docker :)

1. Build 'stats-cli' docker:\
   docker build -t stats-cli .
   
2. Run 'stats-cli' docker as interactive process:\
   docker run -it stats-cli

3. Copy in the standard input the Metadata file files, then press twice Enter to stop the capture

\
Example:\
\
PS C:\Users\benjamin\go\src\github.com\bohana3\stats-cli> docker run -it stats-cli\
2021/08/10 19:06:40 Start stats-cli\
2021/08/10 19:06:40 Enter Metadata files: (one line per file - to stop the capture, press twice Enter!)\
{"path": "C:\\Users\\benjamin\\Downloads\\010EditorWin64Installer801.exe","size": 16594864,"is_binary": true}\
{"path": "C:\\Users\\benjamin\\Downloads\\1513735129.pdf","size": 1552691,"is_binary": false}\
{"path": "C:\\Users\\benjamin\\Downloads\\1513739300.pdf","size": 1552691,"is_binary": false}\
{"path": "C:\\Users\\benjamin\\Downloads\\2017-11 CV template.docx","size": 19484,"is_binary": false}\
{"path": "C:\\Users\\benjamin\\Downloads\\20171215_194251.jpg","size": 2690056,"is_binary": false}\
\
2021/08/10 19:07:32 {\
  "num_files": 5,\
  "largest_file": {\
    "path": "C:\\Users\\benjamin\\Downloads\\010EditorWin64Installer801.exe",\
    "size": 16594864\
  },\
  "avg_file_size": 4481957.2,\
  "most_frequent_ext": {\
    "extension": ".pdf",\
    "num_occurrences": 2\
  },\
  "text_percentage": 80,\
  "most_recent_paths": [\
    "C:\\Users\\benjamin\\Downloads\\010EditorWin64Installer801.exe",\
    "C:\\Users\\benjamin\\Downloads\\1513735129.pdf",\
    "C:\\Users\\benjamin\\Downloads\\1513739300.pdf",\
    "C:\\Users\\benjamin\\Downloads\\2017-11 CV template.docx",\
    "C:\\Users\\benjamin\\Downloads\\20171215_194251.jpg"\
  ]\
}\
2021/08/10 19:07:32 Finish stats-cli
