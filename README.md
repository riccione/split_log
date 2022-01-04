<div id="top"></div>
<h3 align="center">Split log</h3>

<!-- ABOUT THE PROJECT -->
## About The Project

Just a simple cli for Windows. It splits any file by start and end pattern.
<p align="right">(<a href="#top">back to top</a>)</p>



### Built With

* [go](https://go.dev/)

<p align="right">(<a href="#top">back to top</a>)</p>


<!-- USAGE EXAMPLES -->
## Usage

From Powershell or CMD

General example:
>split_log.exe -p="filepath" -s="start" -e="end"

Show on console part of the log from 10:00 to 11:00 on 1/2/2022
>split_log.exe -p="c:\log.txt" -s="1/2/2022 10:00" -e="1/2/2022 11:00"

Show on console part of the log from 10:00 to the end of the file on 1/2/2022
End argument is optional
>split_log.exe -p="c:\log.txt" -s="1/2/2022 10:00"

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- LICENSE -->
## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#top">back to top</a>)</p>