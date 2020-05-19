
# layer2_scanner
## *For when you absolutely need to know whats on your network.*

This is a network scanning tool that relies on the arp protocol and not the TCP/UDP protocols.  Which means it's firewall bypassing, and can tell you if a device is present even if the device won't return a ping or a port discovery request.

### Usage Examples:
> `./layer2_scanner 192.168.1.0/24`


> `./layer2_scanner 192.168.0.1/24`

### Example Output:
>![Example](example.jpg)

