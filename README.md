
# layer2_scanner
## *For when you absolutely need to know whats on your network.*

This is a network scanning tool that relies on the arp protocol and not the TCP/UDP protocols.  Which means it's firewall bypassing, and can tell you if a device is present even if the device won't return a ping or a port discovery request.


---
## Buy me a coffee
If you feel so inclined as to support my projects. Here's your chance! Thanks 

<a href="https://www.buymeacoffee.com/matthewrogers" target="_blank"><img src="https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png" alt="Buy Me A Coffee" style="height: auto !important;width: auto !important;" ></a>

---

**Install**

```
go install github.com/RamboRogers/layer2_scanner/layer2_scanner.go@latest
```

### Usage Examples:
As root or a user with capabilities.
> `./layer2_scanner 192.168.1.0/24`

On newer Macs or Linux you need to escalate.
> `sudo layer2_scanner 192.168.0.1/24`

### Example Output:
>![Example](example.jpg)

#### OS Support:
This application is limited by the libraries that support it. 
1. Linux
2. MacOS
3. BSD

# Homepage / Contact
My homepage and contact info. 

| Type  | Data
| ----- | --------
| Homepage | <https://matthewrogers.org> |

